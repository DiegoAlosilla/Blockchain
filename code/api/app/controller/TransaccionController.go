package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../model"
)

// var direccionesRemotas []Transccion
var transaccionesBD []model.Transaccion

type TransaccionController struct {
}

func (controller *TransaccionController) RegistrarNodo(response http.ResponseWriter, request *http.Request) {
	fmt.Print("Registrar nodo ")
	fmt.Println()
}

func (controller *TransaccionController) RegistrarTransaccion(response http.ResponseWriter, request *http.Request) {
	var transaccion model.Transaccion
	_ = json.NewDecoder(request.Body).Decode(&transaccion)

	transaccionesBD = append(transaccionesBD, transaccion)
	json.NewEncoder(response).Encode(true)

	fmt.Print("new Transaccion: ", transaccion)
	fmt.Println()
}

func (controller *TransaccionController) ListarTransaccionAll(response http.ResponseWriter, request *http.Request) {

	if transaccionesBD == nil {
		vacio := []string{}
		// json.NewEncoder(response).Encode(vacio)
		dataJson, _ := json.Marshal(vacio)
		response.Header().Set("Content-Type", "application/json")
		// response.WriteHeader(http.Status.OK)
		response.Write(dataJson)
	} else {
		dataJson, _ := json.Marshal(transaccionesBD)
		response.Header().Set("Content-Type", "application/json")
		// response.WriteHeader(http.Status.OK)
		response.Write(dataJson)

		// json.NewEncoder(response).Encode(transaccionesBD)
	}

	fmt.Print("transacciones: ", transaccionesBD)
	fmt.Println()
}
