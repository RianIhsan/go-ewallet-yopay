package service

import (
	"errors"
	"fmt"
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/topup"
	"github.com/RianIhsan/go-topup-midtrans/feature/topup/dto"
	"github.com/RianIhsan/go-topup-midtrans/feature/users"
	generator "github.com/RianIhsan/go-topup-midtrans/utils/random"
	"github.com/gofiber/fiber/v2/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type topUpService struct {
	topUpRepo   topup.TopUpRespositoryInterface
	userService users.UserServiceInterface
	generatorID generator.GeneratorInterface
}

func (t topUpService) ConfirmPayment(orderID string) error {
	balance, err := t.topUpRepo.GetBalanceByOrderId(orderID)
	if err != nil {
		return errors.New("transaction data not found")
	}
	balance.Status = "success"

	if err := t.topUpRepo.ConfirmPayment(balance.OrderID, balance.Status); err != nil {
		return err
	}

	getUser, err := t.userService.GetId(balance.UserId)
	if err != nil {
		return errors.New("user not found")
	}
	var totalBalanceAmount float64
	totalBalanceAmount = getUser.TotalBalance + balance.Amount

	err = t.topUpRepo.UpdateUserTotalBalance(balance.UserId, totalBalanceAmount)
	if err != nil {
		return errors.New("failed to update user balance")
	}

	user, err := t.userService.GetId(balance.UserId)
	if err != nil {
		return errors.New("user not found")
	}
	notificationReq := dto.SendNotificationPaymentRequest{
		OrderID:       balance.OrderID,
		UserID:        user.Id,
		PaymentStatus: "success",
	}
	_, err = t.SendNotificationPayment(notificationReq)
	if err != nil {
		log.Error("failed to send notification: ", err)
		return err
	}
	return nil
}

func (t topUpService) SendNotificationPayment(request dto.SendNotificationPaymentRequest) (string, error) {
	var notificationMsg string
	var err error

	user, err := t.userService.GetId(request.UserID)
	if err != nil {
		return "", errors.New("user not found")
	}
	balance, err := t.topUpRepo.GetBalanceByOrderId(request.OrderID)
	if err != nil {
		return "", err
	}
	switch request.PaymentStatus {
	case "pending":
		notificationMsg = fmt.Sprintf("Alloo, %s! Pesananmu dengan ID %s udah berhasil dibuat, nih. Ditunggu yupp!!", user.Fullname, balance.Id)
	case "success":
		notificationMsg = fmt.Sprintf("Thengkyuu, %s! Pembayaran untuk pesananmu dengan ID %s udah kami terima, nih. Semoga harimu menyenangkan!", user.Fullname, balance.Id)
	case "failed":
		notificationMsg = fmt.Sprintf("Maaf, %s. Pembayaran untuk pesanan dengan ID %s gagal, nih. Beritahu kami apabila kamu butuh bantuan yaa!!", user.Fullname, balance.Id)
	default:
		return "", errors.New("status pesanan tidak valid")
	}

	return notificationMsg, nil
}

func (t topUpService) CancelPayment(orderID string) error {
	balance, err := t.topUpRepo.GetBalanceByOrderId(orderID)
	if err != nil {
		return err
	}
	balance.Status = "failed"
	if err := t.topUpRepo.ConfirmPayment(balance.OrderID, balance.Status); err != nil {
		return err
	}
	user, err := t.userService.GetId(balance.UserId)
	if err != nil {
		return errors.New("user not found")
	}
	notificationRequest := dto.SendNotificationPaymentRequest{
		OrderID:       balance.OrderID,
		UserID:        int(user.ID),
		PaymentStatus: "failed",
	}
	_, err = t.SendNotificationPayment(notificationRequest)
	if err != nil {
		log.Error("failed to send notification: ", err)
		return err
	}
	return nil
}

func (t topUpService) CallBack(notifPayload map[string]interface{}) error {
	orderID, exist := notifPayload["order_id"].(string)
	if !exist {
		return errors.New("invalid notification payload")
	}
	status, err := t.topUpRepo.CheckTransaction(orderID)
	if err != nil {
		return err
	}
	transaction, err := t.topUpRepo.GetBalanceByOrderId(orderID)
	if err != nil {
		return errors.New("transaction data not found")
	}
	if status == "success" {
		if err := t.ConfirmPayment(transaction.OrderID); err != nil {
			return err
		}
	} else if status == "failed" {
		if err := t.CancelPayment(transaction.OrderID); err != nil {
			return err
		}
	}
	return nil
}

func (t topUpService) GetTotalBalanceUser(userID int) (float64, error) {
	totalBalance, err := t.topUpRepo.GetBalanceByUserId(userID)
	if err != nil {
		return 0, err
	}
	return totalBalance, nil

}

func (t topUpService) ProcessGatewayPayment(amount int64, orderID string, paymentMethod, fullname, email, phone string, bank midtrans.Bank) (*coreapi.ChargeResponse, error) {
	result, err := t.topUpRepo.ProcessGatewayPayment(amount, orderID, paymentMethod, fullname, email, phone, bank)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t topUpService) CreateTopUp(userId int64, req *dto.TopUpReq, bank midtrans.Bank) (interface{}, error) {
	idOrder, err := t.generatorID.GenerateOrderID()
	if err != nil {
		return nil, errors.New("failed to create order ID")
	}

	if !isValidPaymentMethod(req.PaymentMethod) {
		return nil, errors.New("invalid payment method")
	}

	existingBalance, err := t.topUpRepo.GetBalanceByOrderId(idOrder)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("failed to check existing balance")
	}

	var orderID string

	if existingBalance != nil {
		// Update existing balance
		existingBalance.Amount += float64(req.Amount)
		existingBalance.Status = "pending"
		_, err = t.topUpRepo.UpdateBalance(existingBalance)
		if err != nil {
			return nil, errors.New("failed to update top up data")
		}
		orderID = existingBalance.OrderID
	} else {
		// Create new balance entry
		newBalance := &entities.MstBalance{
			UserId:  int(userId),
			OrderID: idOrder,
			Amount:  float64(req.Amount),
			Status:  "pending",
		}
		_, err = t.topUpRepo.InsertBalance(newBalance)
		if err != nil {
			return nil, errors.New("failed to create top up data")
		}
		orderID = newBalance.OrderID
	}

	user, err := t.userService.GetId(int(userId))
	if err != nil {
		return nil, errors.New("user not found")
	}

	return t.ProcessGatewayPayment(req.Amount, orderID, req.PaymentMethod, user.Fullname, user.Email, user.Phone, bank)
}
func isValidPaymentMethod(method string) bool {
	validPaymentMethods := map[string]bool{
		"qris":          true,
		"bank_transfer": true,
		"gopay":         true,
		"bca":           true,
		"bri":           true,
		"bni":           true,
		"cimb":          true,
	}
	return validPaymentMethods[method]
}

func NewTopUpService(topUpRepo topup.TopUpRespositoryInterface, userService users.UserServiceInterface, generatorId generator.GeneratorInterface) topup.TopUpServiceInterface {
	return &topUpService{
		topUpRepo:   topUpRepo,
		userService: userService,
		generatorID: generatorId,
	}
}
