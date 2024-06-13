package topup

import (
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/topup/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type (
	TopUpRespositoryInterface interface {
		InsertBalance(newData *entities.MstBalance) (*entities.MstBalance, error)
		ProcessGatewayPayment(amount int64, orderID string, paymentMethod, fullname, email, phone string, bank midtrans.Bank) (*coreapi.ChargeResponse, error)
		GetBalanceByUserId(userId int) (float64, error)
		UpdateBalance(balance *entities.MstBalance) (*entities.MstBalance, error)
		GetBalanceById(id string) (*entities.MstBalance, error)
		CheckTransaction(orderID string) (string, error)
		ConfirmPayment(orderID, paymentStatus string) error
		GetBalanceByOrderId(orderID string) (*entities.MstBalance, error)
		UpdateUserTotalBalance(userID int, totalBalance float64) error
		UpdateTotalBalanceByPhone(phone string, totalBalance float64) error
		CreateWithdrawToken(newToken *entities.MstWithdrawBalance) (*entities.MstWithdrawBalance, error)
		GetWithdrawByToken(token int) (*entities.MstWithdrawBalance, error)
		UpdateWithdrawStatus(withdraw *entities.MstWithdrawBalance) error
	}
	TopUpServiceInterface interface {
		CreateTopUp(userId int64, req *dto.TopUpReq, bank midtrans.Bank) (interface{}, error)
		ProcessGatewayPayment(amount int64, orderID string, paymentMethod, fullname, email, phone string, bank midtrans.Bank) (*coreapi.ChargeResponse, error)
		ConfirmPayment(orderID string) error
		SendNotificationPayment(request dto.SendNotificationPaymentRequest) (string, error)
		CancelPayment(orderID string) error
		CallBack(notifPayload map[string]interface{}) error
		GetTotalBalanceUser(userID int) (float64, error)
		TransferBalance(fromUserID int, request dto.TransferBalanceRequest) error
		CreateTokenWithdraw(userId int, req dto.WithdrawBalanceRequest) (*dto.WithdrawBalanceResponse, error)
		GetWithdrawByToken(userId int, req dto.ConfirmWithdrawBalanceRequest) error
	}
	TopUpHandlerInterface interface {
		CreateTopUp(ctx *fiber.Ctx) error
		CallBack(ctx *fiber.Ctx) error
		GetTotalBalance(ctx *fiber.Ctx) error
		TransferBalance(ctx *fiber.Ctx) error
		CreateTokenWithdraw(ctx *fiber.Ctx) error
		ConfirmWithdraw(ctx *fiber.Ctx) error
	}
)
