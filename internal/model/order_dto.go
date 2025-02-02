package model

type OrderDto struct {
	UserId string `json:"user_id" validate:"required"`
}

type UpdateStatusOrderDto struct {
	Status string `json:"status" validate:"required"`
}
