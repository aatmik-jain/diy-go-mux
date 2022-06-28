// main.go

package main

import (
	_ "github.com/lib/pq"
	"go-mux/apis"
	"go-mux/config"
	"log"
	"net/http"
)

func main() {

	db := config.ConnectDB()
	router := apis.Router(db)

	log.Fatal(http.ListenAndServe(":8010", router))
}
