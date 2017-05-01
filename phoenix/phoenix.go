package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/jacadenac/liftit/config"
	"github.com/jacadenac/liftit/services"
	//"github.com/jacadenac/liftit/rabbit"
	"github.com/jacadenac/liftit/rabbit"
	"github.com/jacadenac/liftit/logging"
)

func main(){
	router := mux.NewRouter().StrictSlash(false)
	//defer rabbit.Conn.Close()
	//defer rabbit.Ch.Close()
	services.AddToRouter(router, services.GetUsuarioServ("users"))
	log.Println("Listening http://"+config.Client.Get()+" ...")

	//prueba GETBYID
	for i := 0; i<100; i++{
		usuario, err := rabbit.Publish("GETBYID", []byte(`{"ID":"1"}`))
		logging.FailOnError(err, "Failed GETBYID")
		log.Println("GETBYID = ", usuario)
	}

	http.ListenAndServe(config.Client.Port, router)
}
