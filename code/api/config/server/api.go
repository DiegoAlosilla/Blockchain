package server

import (
	"encoding/json"
	"net/http"

	"../../app/controller"
	"github.com/gorilla/mux"
)

type Server interface {
	Router() http.Handler
}

type Api struct {
	router http.Handler
}

func (api *Api) Router() http.Handler {
	return api.router
}

func InitServer() Server {
	api := &Api{}

	router := mux.NewRouter()

	transaccionController := controller.TransaccionController{}

	router.HandleFunc("/", LinkOfRutas).Methods("GET")

	transaccionRouter := router.PathPrefix("/app").Subrouter()
	transaccionRouter.HandleFunc("/registro", transaccionController.RegistrarNodo).Methods("POST")
	transaccionRouter.HandleFunc("/transacciones", transaccionController.ListarTransaccionAll).Methods("GET")

	api.router = router
	return api
}

func LinkOfRutas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("la ruta general esta en localhost:9000/app")
}
