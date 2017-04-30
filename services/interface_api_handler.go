package services

import (
	"net/http"
	"errors"
	"github.com/gorilla/mux"
	//"../config"
	"log"
	"github.com/jacadenac/liftit/config"
)

var (
	//ErrorFormatoJson ...
	ErrorUsuarioNoValido = errors.New("ErrorFormatoJson")
)

type InterfaceAPIHandler interface {
	getRouteName() string
	getHandler(http.ResponseWriter,*http.Request)
	getByIDHandler(http.ResponseWriter,*http.Request)
	postHandler(http.ResponseWriter,*http.Request)
	putHandler(http.ResponseWriter,*http.Request)
	deleteHandler(http.ResponseWriter,*http.Request)
}

func checkError(err error, w http.ResponseWriter, r *http.Request) bool{
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return true
	}
	return false
}

//AddToRoute - ADD(GET, POST, PUT, DELETE)
func AddToRouter(router *mux.Router, interfaceAPI InterfaceAPIHandler) {
	route := config.Api_root + interfaceAPI.getRouteName()
	router.HandleFunc(route, interfaceAPI.getHandler).Methods("GET")
	router.HandleFunc(route+"/{ID}", interfaceAPI.getByIDHandler).Methods("GET")
	router.HandleFunc(route, interfaceAPI.postHandler).Methods("POST")
	router.HandleFunc(route+"/{ID}", interfaceAPI.putHandler).Methods("PUT")
	router.HandleFunc(route+"/{ID}", interfaceAPI.deleteHandler).Methods("DELETE")
}
