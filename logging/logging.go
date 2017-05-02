package logging

import (
	"log"
	"fmt"
	"encoding/json"
	"net/http"
)

type ServerError struct{
	Error 	string	`json:"error"`
	Detail 	string	`json:"detail"`
}

func (serverError ServerError)ToJson()(json_error []byte) {
	json_error, err := json.Marshal(serverError)
	FailOnError(err, "Error encoding ServerError struct")
	return
}

func JsonError(err error, detail string)([]byte, bool){
	json_error := []byte(`{}`)
	if err != nil {
		serverError := ServerError{err.Error(), detail}
		json_error, err2 := json.Marshal(serverError)
		FailOnError(err2, "Error encoding ServerError struct")
		return json_error, true
	}
	return json_error, false
}

func ResponseError(w http.ResponseWriter, err error, detail string, http_status ...int)(bool){
	if json_error, error := JsonError(err, detail); error {
		if len(http_status) > 0{
			w.WriteHeader(http_status[0])
		}else{
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(json_error)
		return true
	}
	return false
}


func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
