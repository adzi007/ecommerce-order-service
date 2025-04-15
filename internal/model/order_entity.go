package model

type NewOrder struct {
	UserId        string
	TotalPrice    float64
	PaymentMethod string
	PaymentFee    float64
	Status        string
}

type NewOrderDetail struct {
	ID        uint
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
}
