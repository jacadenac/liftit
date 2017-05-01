package controler

import (
	"encoding/json"
	"github.com/jacadenac/liftit/logging"
	"github.com/jacadenac/liftit/contracts"
	"time"
	"strconv"
	"net/http"
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
	type Param struct{
		ID string
	}
	var param Param
	err := json.Unmarshal(request, &param)
	logging.FailOnError(err, "Failed to convert body to json")
	if usuario, ok := contracts.UsuarioStore[param.ID]; ok {
		usuario_privado, err := json.Marshal(usuario)
		var usuario_publico contracts.UsuarioPublico
		err = json.Unmarshal(usuario_privado, &usuario_publico)
		response, err = json.Marshal(usuario_publico)

		logging.FailOnError(err, "Failed to convert usuario to json")
		httpStatus = http.StatusOK
	} else {
		response = []byte(`{}`)
		httpStatus = http.StatusOK
	}
	return
}

func Post(request []byte)(response []byte, httpStatus int) {
	//err := json.Unmarshal(request, contract)
	var usuario contracts.Usuario
	err := json.Unmarshal(request, &usuario)
	logging.FailOnError(err, "Failed to convert json to usuario")
	usuario.CreatedAt = time.Now()
	contracts.ID++
	k := strconv.Itoa(contracts.ID)
	contracts.UsuarioStore[k] = usuario
	response, err = json.Marshal(usuario)
	httpStatus = http.StatusCreated
	return
}

func Put(request []byte)(response []byte, httpStatus int) {
	response = []byte(`{"Put":"usuario editado"}`)
	httpStatus = http.StatusNoContent
	return
}

func Delete(request []byte)(response []byte, httpStatus int) {
	response = []byte(`{"Delete":"usuario borrado"}`)
	httpStatus = http.StatusNoContent
	return
}
