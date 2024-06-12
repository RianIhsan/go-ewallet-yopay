package dto

type TopUpReq struct {
	Amount        int64  `json:"amount" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type TransferBalanceRequest struct {
	Amount float64 `json:"amount" validate:"required"`
	Phone  string  `json:"phone" validate:"required"`
}

type SendNotificationPaymentRequest struct {
	BalanceID     int    `json:"id"`
	PaymentStatus string `json:"payment_status"`
	OrderID       string `json:"order_id"`
	UserID        int    `json:"user_id"`
	Fullname      string `json:"fullname"`
	Title         string `json:"title"`
	Body          string `json:"body"`
}
