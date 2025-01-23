package model

type OrderDto struct {
	UserId string `json:"user_id" validate:"required"`
}
