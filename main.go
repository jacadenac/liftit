package main
/*
import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/jacadenac/liftit/services"
	"github.com/jacadenac/liftit/config"
)

func main(){
	router := mux.NewRouter().StrictSlash(false)
	services.AddToRouter(router, services.GetUsuarioServ("users"))
	log.Println("Listening http://"+config.Server+config.Port+" ...")
	http.ListenAndServe(config.Port, router)
}
*/

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/jacadenac/liftit/phoenix"
	"github.com/jacadenac/liftit/config"
)

func main(){
	router := mux.NewRouter().StrictSlash(false)
	phoenix.AddToRouter(router, phoenix.GetUsuarioServ("users"))
	log.Println("Listening http://"+config.Client.Get()+" ...")
	http.ListenAndServe(config.Client.Port, router)
}
