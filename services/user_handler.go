package services

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"time"
	"strconv"
	"github.com/jacadenac/liftit/config"
	"github.com/jacadenac/liftit/models"
	//"../models"
	//"../config"
	"sync"
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
	usuarios := []models.Usuario{}
	for _, v := range models.UsuarioStore {
		usuarios = append(usuarios, v)
	}
	j, err := json.Marshal(usuarios)
	if !checkError(err, w, r) {
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//getByIDHandler - GET
func (usuario_serv *usuarioServ)getByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	vars := mux.Vars(r)
	ID := vars["ID"]
	if usuario, ok := models.UsuarioStore[ID]; ok {
		j, err := json.Marshal(usuario)
		if !checkError(err, w, r) {
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

//postHandler - POST
func (usuario_serv *usuarioServ)postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	var usuario models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if !checkError(err, w, r) {
		usuario.CreatedAt = time.Now()
		models.ID++
		k := strconv.Itoa(models.ID)
		models.UsuarioStore[k] = usuario
		j, err := json.Marshal(usuario)
		if !checkError(err, w, r){
			w.WriteHeader(http.StatusCreated)
			w.Write(j)
		}
	}
}

//putHandler - PUT
func (usuario_serv *usuarioServ)putHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	vars := mux.Vars(r)
	ID := vars["ID"]
	var usuarioUpdate models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuarioUpdate)
	if !checkError(err, w, r) {
		if usuario, ok := models.UsuarioStore[ID]; ok {
			usuarioUpdate.CreatedAt = usuario.CreatedAt
			delete(models.UsuarioStore, ID)
			models.UsuarioStore[ID] = usuarioUpdate
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

//deleteUsuarioHandler - DELETE
func (usuario_serv *usuarioServ)deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", config.Content_type)
	vars := mux.Vars(r)
	ID := vars["ID"]
	if _, ok := models.UsuarioStore[ID]; ok {
		delete(models.UsuarioStore, ID)
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Println("Not found")
		w.WriteHeader(http.StatusNotFound)
	}
}
