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
	services.AddToRouter(router, services.GetUsuarioServ("users"))
	/*
	//prueba GETBYID
	go func {
		for i := 0; i<100; i++{
			usuario, err := rabbit.Publish("GETBYID", []byte(`{"ID":"1"}`))
			logging.FailOnError(err, "Failed GETBYID")
			log.Println("GETBYID = ", usuario)
		}
	}
	*/
	log.Println("Listening http://"+config.Client.Get()+" ...")
	http.ListenAndServe(config.Client.Port, router)
}
