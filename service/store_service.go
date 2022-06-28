package service

//go:generate mockgen -destination=/Users/aatmikjain/Documents/go-mux/mock/service/store_service.go -source=/Users/aatmikjain/Documents/go-mux/service/store_service.go

import (
	"go-mux/model"
)

type StoreService interface {
	GetProducts(storeID int) ([]model.Product, error)
	AddProducts(storeID int, pIds []int) error
}
