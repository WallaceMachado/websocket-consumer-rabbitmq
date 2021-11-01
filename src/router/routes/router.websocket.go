package routes

import (
	"net/http"

	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/controllers"
)

var routesPerson = []Router{

	{
		URI:    "/",
		Metodo: http.MethodGet,
		Funcao: controllers.HomePage,
	},
	{
		URI:    "/person-api",
		Metodo: http.MethodGet,
		Funcao: controllers.OperationsInPersonApi,
	},
}
