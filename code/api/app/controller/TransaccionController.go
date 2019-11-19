package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../model"
)

var direccionesRemotas []model.Host
var transaccionesBD []model.Transaccion
var Puerto string

type TransaccionController struct {
}

func (controller *TransaccionController) SetPuerto(puerto string) {
	Puerto = puerto
}
func (controller *TransaccionController) MyInfoHost(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(Puerto)
}
func (controller *TransaccionController) TodaLaRed(response http.ResponseWriter, request *http.Request) {
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
func (controller *TransaccionController) NotificandoLaRed(response http.ResponseWriter, request *http.Request) {
	var allHost []model.Host
	_ = json.NewDecoder(request.Body).Decode(&allHost)

	direccionesRemotas = allHost
	json.NewEncoder(response).Encode(direccionesRemotas)
	fmt.Print("Los host en la red son: ", direccionesRemotas)
	fmt.Println()

	// if direccionesRemotas == nil {
	// 	vacio := []string{}
	// 	dataJson, _ := json.Marshal(vacio)
	// 	response.Header().Set("Content-Type", "application/json")
	// 	response.Write(dataJson)
	// } else {
	// 	dataJson, _ := json.Marshal(direccionesRemotas)
	// 	response.Header().Set("Content-Type", "application/json")
	// 	response.Write(dataJson)
	// }
	// fmt.Println()
}

func (controller *TransaccionController) UnirseALaRed(response http.ResponseWriter, request *http.Request) {
	var host model.Host
	_ = json.NewDecoder(request.Body).Decode(&host)

	direccionesRemotas = append(direccionesRemotas, host)

	for _, host := range direccionesRemotas {
		log.Printf("enviar a http://'%s'", host.Ip+":"+host.Puerto)
	}
	// clienteHttp := &http.Client{}
	// url := "http://192.168.0.2:9099/app/unidos"
	// peticion, _ := http.NewRequest("GET", url, nil)
	// respuesta, _ := clienteHttp.Do(peticion)
	// defer respuesta.Body.Close()
	// cuerpoRespuesta, _ := ioutil.ReadAll(respuesta.Body)
	// respuestaString := string(cuerpoRespuesta)
	// log.Printf("Cuerpo de respuesta del servidor: '%s'", respuestaString)

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
