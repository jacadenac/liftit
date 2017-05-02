package controler

import (
	"encoding/json"
	"github.com/jacadenac/liftit/logging"
	"github.com/jacadenac/liftit/contracts"
	"golang.org/x/crypto/bcrypt"
	"time"
	"strconv"
	"net/http"
	"errors"
)

func Get(request []byte)(response []byte, httpStatus int) {
	usuarios := []contracts.Usuario{}
	for _, v := range contracts.UsuarioStore {
		usuarios = append(usuarios, v)
	}
	response, err := json.Marshal(usuarios)
	logging.FailOnError(err, "Failed to convert body to json")
	httpStatus = http.StatusOK
	return
}

func GetById(request []byte)(response []byte, httpStatus int) {
	var payload contracts.Payload
	payload.Set(request)
	if payload.Params == nil{
		response, _ = logging.JsonError(errors.New("Nil Params"), "parametros nulos")
		httpStatus = http.StatusBadRequest
		return
	}
	params := make(map[string]string)
	err := json.Unmarshal(payload.Params, &params)
	if json_error, error := logging.JsonError(err, "Error leyendo parámetros"); error {
		response = json_error
		httpStatus = http.StatusBadRequest
		return
	}
	if params["ID"] == ""{
		response, _ = logging.JsonError(errors.New("Invalid ID"), "Identificador no válido")
		httpStatus = http.StatusBadRequest
		return
	}
	if usuario, ok := contracts.UsuarioStore[params["ID"]]; ok {
		usuario_privado, err := json.Marshal(usuario)
		var usuario_publico contracts.UsuarioPublico
		err = json.Unmarshal(usuario_privado, &usuario_publico)
		response, err = json.Marshal(usuario_publico)
		logging.FailOnError(err, "Failed to convert usuario to json")
		httpStatus = http.StatusOK
	} else {
		response = nil
		httpStatus = http.StatusNotFound
	}
	return
}

func Post(request []byte)(response []byte, httpStatus int) {
	var payload contracts.Payload
	payload.Set(request)
	if payload.Body == nil{
		response, _ = logging.JsonError(errors.New("Nil body"), "cuerpo de petición vacío")
		httpStatus = http.StatusBadRequest
		return
	}
	var usuario contracts.Usuario
	usuario.Set(payload.Body)
	usuario.FechaCreado = time.Now()
	hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	logging.FailOnError(err, "Failed encrypt pass")
	usuario.Password = string(hash[:])

	contracts.ID++
	k := strconv.Itoa(contracts.ID)
	contracts.UsuarioStore[k] = usuario

	response, err = json.Marshal(usuario)
	httpStatus = http.StatusCreated
	return
}

func Put(request []byte)(response []byte, httpStatus int) {
	var payload contracts.Payload
	payload.Set(request)

	// validaciones de parámetros
	if payload.Params == nil{
		response, _ = logging.JsonError(errors.New("Nil Params"), "parametros nulos")
		httpStatus = http.StatusBadRequest
		return
	}
	params := make(map[string]string)
	err := json.Unmarshal(payload.Params, &params)
	if json_error, error := logging.JsonError(err, "Error leyendo parámetros"); error {
		response = json_error
		httpStatus = http.StatusBadRequest
		return
	}
	if params["ID"] == ""{
		response, _ = logging.JsonError(errors.New("Invalid ID"), "Identificador no válido")
		httpStatus = http.StatusBadRequest
		return
	}

	// Validaciones del cuerpo de la petición
	if payload.Body == nil{
		response, _ = logging.JsonError(errors.New("Nil body"), "cuerpo de petición vacío")
		httpStatus = http.StatusBadRequest
		return
	}
	var usuario_edit contracts.Usuario
	usuario_edit.Set(payload.Body)
	hash, err := bcrypt.GenerateFromPassword([]byte(usuario_edit.Password), bcrypt.DefaultCost)
	logging.FailOnError(err, "Failed encrypt pass")
	usuario_edit.Password = string(hash[:])

	// Busca el usuario y lo actualiza
	if usuario, ok := contracts.UsuarioStore[params["ID"]]; ok {
		usuario_edit.FechaCreado = usuario.FechaCreado
		delete(contracts.UsuarioStore, params["ID"])
		contracts.UsuarioStore[params["ID"]] = usuario_edit
		httpStatus = http.StatusNoContent
	} else {
		response = nil
		httpStatus = http.StatusNotFound
	}
	return
}

func Delete(request []byte)(response []byte, httpStatus int) {
	var payload contracts.Payload
	payload.Set(request)
	if payload.Params == nil{
		response, _ = logging.JsonError(errors.New("Nil Params"), "parametros nulos")
		httpStatus = http.StatusBadRequest
		return
	}
	params := make(map[string]string)
	err := json.Unmarshal(payload.Params, &params)
	if json_error, error := logging.JsonError(err, "Error leyendo parámetros"); error {
		response = json_error
		httpStatus = http.StatusBadRequest
		return
	}
	if params["ID"] == ""{
		response, _ = logging.JsonError(errors.New("Invalid ID"), "Identificador no válido")
		httpStatus = http.StatusBadRequest
		return
	}
	if usuario, ok := contracts.UsuarioStore[params["ID"]]; ok {
		delete(contracts.UsuarioStore, params["ID"])
		usuario_json, err := json.Marshal(usuario)
		logging.FailOnError(err, "Failed to convert usuario to json")
		response = usuario_json
		httpStatus = http.StatusNoContent
	} else {
		response = nil
		httpStatus = http.StatusNotFound
	}
	return
}
