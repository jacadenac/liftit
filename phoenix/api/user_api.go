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
	usuarios, err := requestor.Publish("GET", []byte(`{}`))
	//j, err := json.Marshal(usuarios)
	if !checkError(err, w, r) {
		w.WriteHeader(http.StatusOK)
		w.Write(usuarios)
	}
}

//getByIDHandler - GET
func (usuario_serv *usuarioServ)getByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	w.Header().Set(config.Access_control.Key, config.Access_control.Value)
	vars := mux.Vars(r)
	type Param struct{
		ID string
	}
	id_struct := Param{vars["ID"]}
	payload, err := json.Marshal(id_struct)
	//logging.FailOnError(err, "Failed to marshal JSON")
	usuario, err := requestor.Publish("GETBYID", []byte(payload))

	logging.FailOnError(err, "Failed to marshal JSON")

	w.WriteHeader(http.StatusOK)
	w.Write(usuario)

	/*
	if usuario, ok := contracts.UsuarioStore[ID]; ok {
		j, err := json.Marshal(usuario)
		if !checkError(err, w, r) {
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	*/
}

//postHandler - POST
func (usuario_serv *usuarioServ)postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	w.Header().Set(config.Access_control.Key, config.Access_control.Value)
	var usuario contracts.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	log.Println("usuario a crear: ",usuario)
	payload, err := json.Marshal(usuario)
	data, err := requestor.Publish("POST", payload)
	logging.FailOnError(err, "Failed to marshal JSON")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

//putHandler - PUT
func (usuario_serv *usuarioServ)putHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	w.Header().Set(config.Access_control.Key, config.Access_control.Value)
	vars := mux.Vars(r)
	ID := vars["ID"]
	var usuarioUpdate contracts.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuarioUpdate)

	if !checkError(err, w, r) {
		if usuario, ok := contracts.UsuarioStore[ID]; ok {
			usuarioUpdate.CreatedAt = usuario.CreatedAt
			delete(contracts.UsuarioStore, ID)
			contracts.UsuarioStore[ID] = usuarioUpdate
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

//deleteUsuarioHandler - DELETE
func (usuario_serv *usuarioServ)deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	w.Header().Set(config.Access_control.Key, config.Access_control.Value)
	vars := mux.Vars(r)
	ID := vars["ID"]
	if _, ok := contracts.UsuarioStore[ID]; ok {
		delete(contracts.UsuarioStore, ID)
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Println("Not found")
		w.WriteHeader(http.StatusNotFound)
	}
}