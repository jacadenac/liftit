package main

import (

)
import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	//"./services"
	"github.com/jacadenac/liftit/services"
	"github.com/jacadenac/liftit/config"
	//"./config"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)
	services.AddToRouter(router, services.GetUsuarioServ("users"))

	server :=  &http.Server{
		Addr:		config.Server+config.Port,
		Handler:	router,
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
		MaxHeaderBytes:	1 << 20,
	}
	log.Println("Listening http://"+config.Server+config.Port+" ...")
	server.ListenAndServe()
}
