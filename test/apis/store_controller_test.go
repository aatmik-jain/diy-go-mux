package apis

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"go-mux/apis"
	mockService "go-mux/mock/service"
	"go-mux/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*type TestingStruct struct {
	db     *sql.DB
	router *mux.Router
}

var testingStruct TestingStruct

func TestMain(m *testing.M) {
	testingStruct.db = config.ConnectDB()
	testingStruct.router = apis.Router(testingStruct.db)
	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := testingStruct.db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := testingStruct.db.Exec(storeTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	testingStruct.db.Exec("DELETE FROM products")
	testingStruct.db.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")

	testingStruct.db.Exec("DELETE FROM store")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
   id SERIAL,
   name TEXT NOT NULL,
   price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
   CONSTRAINT products_pkey PRIMARY KEY (id)
)`

const storeTableCreationQuery = `CREATE TABLE IF NOT EXISTS store
(
	store_id SERIAL PRIMARY KEY,
	product_id integer,
	is_available BOOLEAN,
	FOREIGN KEY(product_id) REFERENCES products(id)
)`

func TestEmptyTable(t *testing.T) {
	clearTable()

	//req, _ := http.NewRequest("GET", "/products", nil)
	//response := executeRequest(req)
	//
	//checkResponseCode(t, http.StatusOK, response.Code)
	//
	//if body := response.Body.String(); body != "[]" {
	//	t.Errorf("Expected an empty array. Got %s", body)
	//}

	// null stores
	req, _ := http.NewRequest("GET", "/store/2/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "null" {
		t.Errorf("Expected null. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	testingStruct.router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestAddProductsToStore(t *testing.T) {
	clearTable()

	var body = []byte(`{"productIDs":[1,2,4]}`)
	req, _ := http.NewRequest("POST", "/store/10", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var resString string
	json.Unmarshal(response.Body.Bytes(), &resString)

	if resString != "Product added successfully" {
		t.Errorf("Expected store id to be \"Product(s) added successfully\". Got '%v'", resString)
	}
}

func TestGetProductsFromStore(t *testing.T) {
	clearTable()
	addProducts(10, []int{1, 2, 3, 4, 5})

	req, _ := http.NewRequest("GET", "/store/10/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
func addProducts(storeID int, pIDs []int) {

	for _, id := range pIDs {
		testingStruct.db.QueryRow(
			"INSERT INTO store(store_id, product_id, is_available) VALUES($1, $2, true)",
			storeID, id)
	}
}*/

func TestAddProductToStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	//InputProduct := &dtos.Product{Name: "Shirt", Description: "Denim Fabric", Price: 2499, Quantity: 582}

	body := []byte(`{"productIDs":[1,2,4]}`)

	req, _ := http.NewRequest("POST", "/store/11", bytes.NewBuffer(body))
	mockServices := mockService.NewMockStoreService(ctrl)
	mockServices.EXPECT().AddProducts(11, []int{1, 2, 4}).Return(nil)
	storeController := apis.StoreController{StoreService: mockServices}

	r.HandleFunc("/store/{id:[0-9]+}", storeController.AddProducts).Methods("POST")
	r.ServeHTTP(rr, req)

	stringResponse := `"Product(s) added successfully to store: 11"`
	str := strings.TrimSpace(rr.Body.String())
	assert.Equal(t, stringResponse, str)
	assert.Equal(t, http.StatusCreated, rr.Code)

}

func TestGetProductsFromStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	output := []model.Product{{ID: 1, Name: "prod1", Price: 10}, {ID: 4, Name: "prod4", Price: 40}}

	body := []byte(``)

	req, _ := http.NewRequest("GET", "/store/10/products", bytes.NewBuffer(body))
	mockServices := mockService.NewMockStoreService(ctrl)
	mockServices.EXPECT().GetProducts(10).Return(output, nil)
	storeController := apis.StoreController{StoreService: mockServices}

	r.HandleFunc("/store/{id:[0-9]+}/products", storeController.GetProducts).Methods("GET")
	r.ServeHTTP(rr, req)

	stringResponse := `[{"id":1,"name":"prod1","price":10},{"id":4,"name":"prod4","price":40}]`
	str := strings.TrimSpace(rr.Body.String())
	assert.Equal(t, stringResponse, str)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetProductsFromStoreNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	output := []model.Product{{ID: 1, Name: "prod1", Price: 10}, {ID: 4, Name: "prod4", Price: 40}}

	body := []byte(``)

	req, _ := http.NewRequest("GET", "/store/25/products", bytes.NewBuffer(body))
	mockServices := mockService.NewMockStoreService(ctrl)
	mockServices.EXPECT().GetProducts(25).Return(output, errors.New("Store not found"))
	storeController := apis.StoreController{StoreService: mockServices}

	r.HandleFunc("/store/{id:[0-9]+}/products", storeController.GetProducts).Methods("GET")
	r.ServeHTTP(rr, req)

	stringResponse := `{"error":"Store not found"}`
	str := strings.TrimSpace(rr.Body.String())
	assert.Equal(t, stringResponse, str)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

//go test -covermode=count  -p 1  -coverpkg=../main/... ./... -coverprofile=cover.out && go tool cover -func=cover.out
