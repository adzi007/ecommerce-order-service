package domain

type CartItem struct {
	ID        uint64
	ProductId uint64
	Name      string
	Slug      string
	Price     uint64
	Qty       uint64
	Category  ProductCategory
}

type ProductCategory struct {
	Name string
	Slug string
}

type DeleteCartResponse struct {
	Message string
}

type CartRepository interface {
	GetCartByUserID(userID string) ([]CartItem, error)
	DeleteCartUser(userID string) (DeleteCartResponse, error)
}
