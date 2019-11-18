package controller

import (
	"fmt"
	"net/http"
)

type TransaccionController struct {
}

func (controller *TransaccionController) RegistrarNodo(response http.ResponseWriter, request *http.Request) {
	fmt.Print("Registrar nodo ")
}

func (controller *TransaccionController) RegistrarTransaccion(response http.ResponseWriter, request *http.Request) {
	fmt.Print("Registrar transaccion ")
}

func (controller *TransaccionController) ListarTransaccionAll(response http.ResponseWriter, request *http.Request) {
	fmt.Print("Listar transacciones ")
}
