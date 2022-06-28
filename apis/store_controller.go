package apis

//go:generate mockgen -destination=/Users/aatmikjain/Documents/go-mux/mock/apis/store_controller.go -source=/Users/aatmikjain/Documents/go-mux/apis/store_controller.go

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-mux/service"
	"go-mux/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

type StoreController struct {
	StoreService service.StoreService
}

type StoreInterface interface {
	GetProductsFromStore(w http.ResponseWriter, r *http.Request)
	AddProductsToStore(w http.ResponseWriter, r *http.Request)
}

func (scs StoreController) GetProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	storeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid Store ID")
		return
	}
	products, err := scs.StoreService.GetProducts(storeID)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Store not found")
		return
	}

	util.RespondWithJSON(w, http.StatusOK, products)
}

func (scs StoreController) AddProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid Store ID")
	}

	type productList struct {
		ProductIDs []int
	}
	var pl productList
	body, err := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &pl); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := scs.StoreService.AddProducts(storeID, pl.ProductIDs); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, "Product(s) added successfully to store: "+strconv.Itoa(storeID))
}
