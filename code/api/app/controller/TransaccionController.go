package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../model"
)

var direccionesRemotas []model.Host
var transaccionesBD []model.Transaccion
var Puerto int

type TransaccionController struct {
	// puerto int
}

// func (controller *TransaccionController) SetPuerto(puerto int) {
// 	// controller.puerto = puerto
// }
func (controller *TransaccionController) MyInfoHost(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("puerto")
}
func (controller *TransaccionController) ListarHostConectados(response http.ResponseWriter, request *http.Request) {

	if direccionesRemotas == nil {
		vacio := []string{}
		dataJson, _ := json.Marshal(vacio)
		response.Header().Set("Content-Type", "application/json")
		response.Write(dataJson)
	} else {
		dataJson, _ := json.Marshal(direccionesRemotas)
		response.Header().Set("Content-Type", "application/json")
		response.Write(dataJson)
	}
	fmt.Println()
}

func (controller *TransaccionController) ConectarUnHost(response http.ResponseWriter, request *http.Request) {
	var host model.Host
	_ = json.NewDecoder(request.Body).Decode(&host)

	direccionesRemotas = append(direccionesRemotas, host)
	json.NewEncoder(response).Encode(true)

	fmt.Print("Nodo Agregado Exitosamente: ", host)
	fmt.Println()
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

	fmt.Print("Transaccion Agregado Exitosamente: ", transaccion)
	fmt.Println()
}

func (controller *TransaccionController) ListarTransaccionAll(response http.ResponseWriter, request *http.Request) {

	if transaccionesBD == nil {
		vacio := []string{}
		// json.NewEncoder(response).Encode(vacio)
		dataJson, _ := json.Marshal(vacio)
		response.Header().Set("Content-Type", "application/json")
		response.Write(dataJson)
	} else {
		dataJson, _ := json.Marshal(transaccionesBD)
		response.Header().Set("Content-Type", "application/json")
		response.Write(dataJson)
		// json.NewEncoder(response).Encode(transaccionesBD)
	}

	fmt.Print("transacciones: ", transaccionesBD)
	fmt.Println()
}
