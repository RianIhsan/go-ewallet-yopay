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

type WithdrawBalanceRequest struct {
	Amount   float64 `json:"amount" validate:"required"`
	Provider string  `json:"provider" validate:"required"`
}

type WithdrawBalanceResponse struct {
	Token        int   `json:"token"`
	TokenExpired int64 `json:"token_expired"`
}

type ConfirmWithdrawBalanceRequest struct {
	Token int `json:"token" validate:"required"`
}
