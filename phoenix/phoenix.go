package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/jacadenac/liftit/config"
	"github.com/jacadenac/liftit/services"
	//"github.com/jacadenac/liftit/rabbit"
)

func main(){
	router := mux.NewRouter().StrictSlash(false)
	//defer rabbit.Conn.Close()
	//defer rabbit.Ch.Close()
	services.AddToRouter(router, services.GetUsuarioServ("users"))
	log.Println("Listening http://"+config.Client.Get()+" ...")
	http.ListenAndServe(config.Client.Port, router)
}
