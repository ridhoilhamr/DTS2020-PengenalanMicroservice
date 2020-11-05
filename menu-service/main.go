package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ridhoilhamr/DTS2020-PengenalanMicroservice/menu-service/handler"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Menu service listen on port :5000")
	log.Panic(http.ListenAndServe(":5000", router))
}
