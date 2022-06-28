package apis

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go-mux/service/impl"
)

var router *mux.Router

func Router(db *sql.DB) *mux.Router {
	router = mux.NewRouter()

	initializeRoutes(db)

	return router
}

func initializeRoutes(db *sql.DB) {

	storeService := impl.StoreServiceImpl{DB: db}
	storeController := StoreController{StoreService: storeService}

	router.HandleFunc("/store/{id:[0-9]+}/products", storeController.GetProducts).Methods("GET")
	router.HandleFunc("/store/{id:[0-9]+}", storeController.AddProducts).Methods("POST")
}
