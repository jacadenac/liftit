package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/jacadenac/liftit/services"
	"github.com/jacadenac/liftit/config"
)

func main(){
	Port := ":9000"
	router := mux.NewRouter().StrictSlash(false)
	services.AddToRouter(router, services.GetUsuarioServ("users"))
	log.Println("Listening http://"+config.Server+Port+" ...")
	http.ListenAndServe(Port, router)
}
