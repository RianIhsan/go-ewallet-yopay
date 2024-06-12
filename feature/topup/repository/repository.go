package repository

import (
	"errors"
	"fmt"
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/topup"
	"github.com/RianIhsan/go-topup-midtrans/utils/payment"
	"github.com/gofiber/fiber/v2/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type topUpRepo struct {
	db         *gorm.DB
	coreClient coreapi.Client
}

func (t *topUpRepo) GetBalanceById(id string) (*entities.MstBalance, error) {
	var balance entities.MstBalance
	if err := t.db.Where("id = ? ", id).First(&balance).Error; err != nil {
		fmt.Println("ERRORRR Get balance By Id")
		return nil, err
	}
	return &balance, nil
}

func (t *topUpRepo) GetBalanceByOrderId(orderID string) (*entities.MstBalance, error) {
	var balance entities.MstBalance
	if err := t.db.Where("order_id = ? ", orderID).First(&balance).Error; err != nil {
		return nil, err
	}
	return &balance, nil
}

func (t *topUpRepo) CheckTransaction(orderID string) (string, error) {
	var paymentStatus string
	transactionStatusResp, err := t.coreClient.CheckTransaction(orderID)
	if err != nil {
		return "", err
	} else {
		if transactionStatusResp != nil {
			paymentStatus = payment.TransactionStatus(transactionStatusResp)
			return paymentStatus, nil
		}
	}
	return "", errors.New("transaction not found")
}

func (t *topUpRepo) ConfirmPayment(orderID, paymentStatus string) error {
	var balance entities.MstBalance
	if err := t.db.Model(&balance).Where("order_id = ?", orderID).Updates(map[string]interface{}{
		"status": paymentStatus,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (t *topUpRepo) InsertBalance(newData *entities.MstBalance) (*entities.MstBalance, error) {
	err := t.db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (t *topUpRepo) ProcessGatewayPayment(amount int64, orderID string, paymentMethod, fullname, email, phone string, bank midtrans.Bank) (*coreapi.ChargeResponse, error) {
	var paymentType coreapi.CoreapiPaymentType

	switch paymentMethod {
	case "qris":
		paymentType = coreapi.PaymentTypeQris
	case "bri", "bni", "bca", "cimb", "permata":
		paymentType = coreapi.PaymentTypeBankTransfer
	case "gopay":
		paymentType = coreapi.PaymentTypeGopay
	}
	coreClient := t.coreClient
	resp, err := payment.CreateCoreAPIPaymentRequest(coreClient, orderID, amount, paymentType, fullname, email, phone, bank)
	if err != nil {
		log.Error(err)
		return nil, errors.New("failed to create payment request")
	}
	return resp, nil
}

func (t *topUpRepo) GetBalanceByUserId(userId int) (float64, error) {
	var totalBalance float64
	err := t.db.Model(&entities.MstUser{}).
		Where("id = ?", userId).
		Select("total_balance").
		Row().
		Scan(&totalBalance)
	if err != nil {
		return 0, err
	}
	return totalBalance, nil
}

func (t *topUpRepo) UpdateBalance(balance *entities.MstBalance) (*entities.MstBalance, error) {
	err := t.db.Save(balance).Error
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (t *topUpRepo) UpdateUserTotalBalance(userID int, totalBalance float64) error {
	return t.db.Model(&entities.MstUser{}).
		Where("id = ?", userID).
		Update("total_balance", totalBalance).
		Error
}

func (t *topUpRepo) UpdateTotalBalanceByPhone(phone string, totalBalance float64) error {
	if err := t.db.Model(&entities.MstUser{}).
		Where("phone = ? AND deleted_at is NULL", phone).
		Update("total_balance", totalBalance).
		Error; err != nil {
		return err
	}
	return nil
}

func NewTopUpRepository(db *gorm.DB, coreClient coreapi.Client) topup.TopUpRespositoryInterface {
	return &topUpRepo{
		db:         db,
		coreClient: coreClient,
	}
}
