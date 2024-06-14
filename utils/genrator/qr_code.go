package generator

import (
	"context"
	"fmt"
	"github.com/RianIhsan/go-topup-midtrans/entities"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/skip2/go-qrcode"
)

type QrCodeInterface interface {
	GenerateQRCode(user *entities.MstUser) error
}

type qrCode struct {
	cloudinary *cloudinary.Cloudinary
}

func NewQrGenerator(cloudinary *cloudinary.Cloudinary) QrCodeInterface {
	return &qrCode{
		cloudinary: cloudinary,
	}
}

func (q qrCode) GenerateQRCode(user *entities.MstUser) error {
	qrData := user.Phone
	qrCode, err := qrcode.New(qrData, qrcode.Medium)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s.png", user.Phone)
	filePath := fmt.Sprintf("/tmp/%s", fileName)

	err = qrCode.WriteFile(256, filePath)
	if err != nil {
		return err
	}

	ctx := context.Background()
	uploadResult, err := q.cloudinary.Upload.Upload(ctx, filePath, uploader.UploadParams{PublicID: fileName})
	if err != nil {
		return err
	}

	// Set URL QR code
	user.QRCode = uploadResult.SecureURL
	return nil
}
