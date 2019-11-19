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

var pu string

func InitServer(puerto string) Server {
	api := &Api{}
	pu = puerto
	router := mux.NewRouter()

	transaccionController := controller.TransaccionController{}
	// transaccionController.SetPuerto(puerto)

	router.HandleFunc("/", LinkOfRutas).Methods("GET")

	transaccionRouter := router.PathPrefix("/app").Subrouter()
	transaccionRouter.HandleFunc("/registro", transaccionController.RegistrarNodo).Methods("POST")
	transaccionRouter.HandleFunc("/transaccion", transaccionController.RegistrarTransaccion).Methods("POST")
	transaccionRouter.HandleFunc("/transacciones", transaccionController.ListarTransaccionAll).Methods("GET")

	transaccionRouter.HandleFunc("/host", transaccionController.ListarHostConectados).Methods("GET")
	transaccionRouter.HandleFunc("/host", transaccionController.ConectarUnHost).Methods("POST")
	transaccionRouter.HandleFunc("/myhost", transaccionController.MyInfoHost).Methods("GET")
	api.router = router
	return api
}

func LinkOfRutas(w http.ResponseWriter, r *http.Request) {
	// i := 123
	// p := strconv.Itoa(pu)
	result := "la ruta general esta en localhost:" + pu + "/app"
	json.NewEncoder(w).Encode(result)
}
