package service

import "github.com/adzi007/ecommerce-order-service/internal/domain"

type CartUsecase struct {
	cartRepo domain.CartRepository
}

func NewCartUsecase(repo domain.CartRepository) *CartUsecase {
	return &CartUsecase{cartRepo: repo}
}

func (u *CartUsecase) GetCartByUserID(userID string) ([]domain.CartItem, error) {
	return u.cartRepo.GetCartByUserID(userID)
}
