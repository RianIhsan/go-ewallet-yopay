package handler

import (
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/RianIhsan/go-topup-midtrans/feature/topup"
	"github.com/RianIhsan/go-topup-midtrans/feature/topup/dto"
	"github.com/RianIhsan/go-topup-midtrans/utils/response"
	"github.com/RianIhsan/go-topup-midtrans/utils/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
)

type topUpHandler struct {
	topUpService topup.TopUpServiceInterface
}

func (t *topUpHandler) GetTotalBalance(ctx *fiber.Ctx) error {
	currentUser, _ := ctx.Locals("CurrentUser").(*entities.MstUser)
	if currentUser == nil {
		return response.SendStatusUnauthorized(ctx, "access denied")
	}
	totalBalance, err := t.topUpService.GetTotalBalanceUser(currentUser.Id)
	if err != nil {
		return response.SendStatusInternalServerError(ctx, "failed to get total balance")
	}
	return response.SendStatusOkWithDataResponse(ctx, "success", dto.GetTotalBalanceUser(currentUser, totalBalance))
}

func (t *topUpHandler) CallBack(ctx *fiber.Ctx) error {
	var notificationPayload map[string]any

	if err := ctx.BodyParser(&notificationPayload); err != nil {
		return response.SendStatusBadRequest(ctx, "invalid request")
	}
	err := t.topUpService.CallBack(notificationPayload)
	if err != nil {
		return response.SendStatusInternalServerError(ctx, "failed to process notification")
	}
	return response.SendStatusOkResponse(ctx, "success callback")
}

func (t *topUpHandler) CreateTopUp(ctx *fiber.Ctx) error {
	currentUser, _ := ctx.Locals("CurrentUser").(*entities.MstUser)
	if currentUser == nil {
		return response.SendStatusUnauthorized(ctx, "access denied")
	}

	var payload *dto.TopUpReq
	if err := ctx.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(ctx, "invalid request"+err.Error())
	}
	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(ctx, "error validating payload:"+err.Error())
	}

	var bank midtrans.Bank
	switch payload.PaymentMethod {
	case "bca":
		bank = midtrans.BankBca
	case "bri":
		bank = midtrans.BankBri
	case "bni":
		bank = midtrans.BankBni
	case "cimb":
		bank = midtrans.BankCimb
	default:
		bank = midtrans.BankPermata
	}
	result, err := t.topUpService.CreateTopUp(int64(currentUser.Id), payload, bank)
	if err != nil {
		return response.SendStatusBadRequest(ctx, "failed to create top up data:"+err.Error())
	}
	switch payload.PaymentMethod {
	case "qris", "bank_transfer", "bca", "bri", "bni", "cimb", "gopay":
		return response.SendStatusCreatedWithDataResponse(ctx, "Berhasil membuat pesanan dengan payment gateway", result)
	default:
		return response.SendStatusBadRequest(ctx, "Metode pembayaran tidak valid: "+payload.PaymentMethod)
	}
}

func (t *topUpHandler) TransferBalance(ctx *fiber.Ctx) error {
	currentUser, _ := ctx.Locals("CurrentUser").(*entities.MstUser)
	if currentUser == nil {
		return response.SendStatusUnauthorized(ctx, "access denied")
	}

	var payload dto.TransferBalanceRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(ctx, "invalid request")
	}
	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(ctx, "error validating payload")
	}
	err := t.topUpService.TransferBalance(currentUser.Id, payload)
	if err != nil {
		return response.SendStatusBadRequest(ctx, "failed to transfer balance: "+err.Error())
	}
	return response.SendStatusOkResponse(ctx, "success transfer balance")
}

func NewTopUpHandler(topUpService topup.TopUpServiceInterface) topup.TopUpHandlerInterface {
	return &topUpHandler{
		topUpService: topUpService,
	}
}
