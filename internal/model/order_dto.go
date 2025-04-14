package model

type OrderDto struct {
	UserId        string  `json:"user_id" validate:"required"`
	PaymentMethod string  `json:"payment_method" validate:"required"`
	PaymentFee    float64 `json:"payment_fee" validate:"required"`
}

type UpdateStatusOrderDto struct {
	Status string `json:"status" validate:"required"`
}
