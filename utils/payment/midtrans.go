package payment

import (
	"github.com/RianIhsan/go-topup-midtrans/config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func InitMidtransCore(config config.Config) coreapi.Client {
	var coreClient coreapi.Client
	coreClient.New(config.Midtrans.ServerKey, midtrans.Sandbox)
	return coreClient
}
