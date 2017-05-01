package logging

import (
	"log"
	"fmt"
	"encoding/json"
)

type ServerError struct{
	Error 	error	`json:"error"`
	Detail 	string	`json:"detail"`
}

func (serverError ServerError)toJson(json_error []byte) {
	json_error, err := json.Marshal(serverError)
	FailOnError(err, "Error encoding ServerError struct")
	return
}

func JsonError(err error, detail string)([]byte, bool){
	json_error := []byte(`{}`)
	if err != nil {
		serverError := ServerError{err, detail}
		json_error, err2 := json.Marshal(serverError)
		FailOnError(err2, "Error encoding ServerError struct")
		return json_error, true
	}
	return json_error, false
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
