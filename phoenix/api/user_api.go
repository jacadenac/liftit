package api

import (
	"log"
	"sync"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jacadenac/liftit/config"
	"github.com/jacadenac/liftit/contracts"
	"github.com/jacadenac/liftit/logging"
	"github.com/jacadenac/liftit/logging/detail"
	"github.com/jacadenac/liftit/phoenix/requestor"
)

//Se implementa estructura usuarioServ bajo patron singleton
type usuarioServ struct{
	Name_route string
}
var instance *usuarioServ
var once sync.Once
func GetUsuarioServ(name_route string) *usuarioServ {
	once.Do(func() {
		instance = &usuarioServ{name_route}
	})
	return instance
}

//usuarioServe debe implemetar los siguientes métodos de la interfaz InterfaceAPIHandler:
//getRouteName() string
//getHandler(http.ResponseWriter,*http.Request)
//getByIDHandler(http.ResponseWriter,*http.Request)
//postHandler(http.ResponseWriter,*http.Request)
//putHandler(http.ResponseWriter,*http.Request)
//deleteHandler(http.ResponseWriter,*http.Request)

//getRouteName
func (usuario_serv *usuarioServ)getRouteName() string{
	return usuario_serv.Name_route
}

//getHandler - GET
func (usuario_serv *usuarioServ)getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	w.Header().Set(config.Access_control.Key, config.Access_control.Value)
	payload := contracts.Payload{nil, nil}
	usuarios, err := requestor.Publish("GET", payload.ToJson())
	if logging.ResponseError(w, err, detail.FailedPublish){
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(usuarios)
}

//getByIDHandler - GET
func (usuario_serv *usuarioServ)getByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	w.Header().Set(config.Access_control.Key, config.Access_control.Value)
	vars, err := json.Marshal(mux.Vars(r))
	if logging.ResponseError(w, err, detail.FailedParameterReading, http.StatusBadRequest){
		return
	}
	payload := contracts.Payload{nil, vars}
	data, err := requestor.Publish("GETBYID", payload.ToJson())
	if logging.ResponseError(w, err, detail.FailedPublish){
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//postHandler - POST
func (usuario_serv *usuarioServ)postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	w.Header().Set(config.Access_control.Key, config.Access_control.Value)
	var usuario contracts.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if logging.ResponseError(w, err, detail.FailedFormatJson, http.StatusBadRequest){
		return
	}
	body, err := json.Marshal(usuario)
	if logging.ResponseError(w, err, detail.FailedConvertToJson, http.StatusBadRequest){
		return
	}
	payload := contracts.Payload{body, nil}
	log.Println("usuario a crear: ",payload.ToJson())
	data, err := requestor.Publish("POST", payload.ToJson())
	if logging.ResponseError(w, err, detail.FailedPublish){
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

//putHandler - PUT
func (usuario_serv *usuarioServ)putHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	w.Header().Set(config.Access_control.Key, config.Access_control.Value)
	vars, err := json.Marshal(mux.Vars(r))
	if logging.ResponseError(w, err, detail.FailedParameterReading, http.StatusBadRequest){
		return
	}
	var usuario contracts.Usuario
	err = json.NewDecoder(r.Body).Decode(&usuario)
	if logging.ResponseError(w, err, detail.FailedFormatJson, http.StatusBadRequest){
		return
	}
	body, err := json.Marshal(usuario)
	payload := contracts.Payload{body, vars}
	data, err := requestor.Publish("PUT", payload.ToJson())
	if logging.ResponseError(w, err, detail.FailedPublish){
		return
	}
	if data != nil{
		w.WriteHeader(http.StatusNoContent)
	}else{
		w.WriteHeader(http.StatusNotFound)
	}
}

//deleteUsuarioHandler - DELETE
func (usuario_serv *usuarioServ)deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	w.Header().Set(config.Access_control.Key, config.Access_control.Value)
	vars, err := json.Marshal(mux.Vars(r))
	if logging.ResponseError(w, err, detail.FailedParameterReading, http.StatusBadRequest){
		return
	}
	payload := contracts.Payload{nil, vars}
	data, err := requestor.Publish("DELETE", payload.ToJson())
	if logging.ResponseError(w, err, detail.FailedPublish){
		return
	}
	if data != nil{
		w.WriteHeader(http.StatusNoContent)
	}else{
		w.WriteHeader(http.StatusNotFound)
	}
}
