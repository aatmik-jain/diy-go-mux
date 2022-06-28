package model

type Store struct {
	StoreId     int  `json:"store_id"`
	ProductId   int  `json:"product_id"`
	IsAvailable bool `json:"is_available"`
}
