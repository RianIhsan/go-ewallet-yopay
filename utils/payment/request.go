package payment

import (
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CreateCoreAPIPaymentRequest(
	coreClient coreapi.Client,
	orderID string,
	amount int64,
	paymentType coreapi.CoreapiPaymentType,
	fullname string,
	email string,
	phone string,
	bank midtrans.Bank,
) (*coreapi.ChargeResponse, error) {
	var PaymentRequest *coreapi.ChargeReq

	switch paymentType {
	case coreapi.PaymentTypeQris, coreapi.PaymentTypeGopay, coreapi.PaymentTypeBankTransfer:
		PaymentRequest = &coreapi.ChargeReq{
			PaymentType: paymentType,
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderID,
				GrossAmt: amount,
			},
		}
		if paymentType == coreapi.PaymentTypeBankTransfer {
			PaymentRequest.BankTransfer = &coreapi.BankTransferDetails{
				Bank: bank,
			}
		}
	default:
		return nil, errors.New("payment type not supported")
	}

	PaymentRequest.CustomerDetails = &midtrans.CustomerDetails{
		FName: fullname,
		Phone: phone,
		Email: email,
	}
	resp, err := coreClient.ChargeTransaction(PaymentRequest)
	if err != nil {
		log.Error("Error creating payment request:", err.Error())
		return nil, err
	}
	log.Info("save payment data: OrderId = ", orderID, "fullname=", fullname, "Email=", email, "Phone=", phone)
	log.Info("Payment request created successfully:", resp)
	return resp, nil
}
