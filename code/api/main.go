package main

import (
	"log"
	"net/http"

	"./config/server"
)

func main() {
	api := server.InitServer()
	log.Fatal(http.ListenAndServe(":9000", api.Router()))
}
