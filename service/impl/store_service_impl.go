package impl

import (
	"database/sql"
	"go-mux/model"
)

type StoreServiceImpl struct {
	DB *sql.DB
}

func (s StoreServiceImpl) GetProducts(storeID int) ([]model.Product, error) {

	//TODO: try with redis

	rows, err := s.DB.Query(
		"SELECT COUNT(*) FROM store WHERE store_id=$1",
		storeID)
	if err != nil {
		return nil, err
	}
	var count int
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return nil, err
		}
	}
	if count <= 0 {
		err := sql.ErrNoRows
		return nil, err
	}

	rows, err = s.DB.Query(
		"SELECT products.id, products.name, products.price FROM store, products WHERE store.store_id=$1 AND store.product_id=products.id AND store.is_available",
		storeID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (s StoreServiceImpl) AddProducts(storeID int, pIds []int) error {
	// 1. NO FOREIGN KEY not checking of product id is present in product table
	// meaning - there might be a product present in store table, but not in product table
	// 2. NO PRIMARY KEY - MAY CONTAIN DUPLICATE ROWS

	var err error
	for _, pid := range pIds {
		err = s.DB.QueryRow(
			"INSERT INTO store(store_id, product_id, is_available) VALUES($1, $2, true) RETURNING product_id",
			storeID, pid).Scan(&pid)
	}
	if err != nil {
		return err
	}

	return nil
}
