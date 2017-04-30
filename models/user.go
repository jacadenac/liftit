package models

import (
	"time"
	"strconv"
)

var UsuarioStore = make(map[string]Usuario)
var ID int

type Usuario struct {
	Nombre		string		`json:"nombre"`
	Email		string		`json:"description"`
	Password 	string		`json:"password"`
	Verificado	bool		`json:"verificado"`
	Telefono	int		`json:"telefono"`
	Pais		string		`json:"pais"`
	Ciudad		string		`json:"Ciudad"`
	Direccion	string		`json:"Direccion"`
	CreatedAt 	time.Time 	`json:"created_at"`
}

func init(){
	usuario := Usuario{
		"Juan David Alejandro Cadena Cedano",
		"jacadenac@unal.edu.co",
		"juan14",
		true,
		3016598028,
		"Colombia",
		"Bogotá",
		"Cll 19 B sur No. 53A - 28",
		time.Now(),
	}
	ID++
	k := strconv.Itoa(ID)
	UsuarioStore[k] = usuario
}