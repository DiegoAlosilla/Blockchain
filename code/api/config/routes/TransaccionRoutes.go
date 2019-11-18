package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type TransaccionRoutes struct {
	MyRouter http.Handler
}

func (transaccion *TransaccionRoutes) Routes() {
	transaccion.MyRouter = mux.NewRouter()

	// poki := pokemon.MyRouter.PathPrefix("/pokemon2").Subrouter()
	// poki.HandleFunc("/2", Insertar1).Methods("GET")
	// router := mux.NewRouter()
	// // router := pokemon.myRouter
	// pokeRoutes := router.PathPrefix("/pokemon2").Subrouter()

	// pokeRoutes.HandleFunc("/2", Insertar1).Methods("GET")
	// pokeRoutes.HandleFunc("/3", Insertar2).Methods("GET")
}
