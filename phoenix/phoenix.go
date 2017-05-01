package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/jacadenac/liftit/config"
	"github.com/jacadenac/liftit/phoenix/api"
)



func main(){
	router := mux.NewRouter().StrictSlash(false)
	api.AddToRouter(router, api.GetUsuarioServ("users"))
	log.Println("Listening http://"+config.Client.Get()+" ...")
	http.ListenAndServe(config.Client.Port, router)
}
