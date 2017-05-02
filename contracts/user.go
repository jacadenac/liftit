package contracts

import (
	"time"
	"strconv"
	//"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"github.com/jacadenac/liftit/logging"
)

var UsuarioStore = make(map[string]Usuario)
var ID int

type Usuario struct {
	Nombre		string		`json:"nombre, required"`
	Email		string		`json:"email" binding:"required"`
	Password 	string		`json:"password" binding:"required"`
	Verificado	bool		`json:"verificado, omitempty"`
	Telefono	int		`json:"telefono" binding:"required"`
	Pais		string		`json:"pais, omitempty"`
	Ciudad		string		`json:"ciudad, omitempty"`
	Direccion	string		`json:"direccion, omitempty"`
	FechaCreado	time.Time 	`json:"fecha_creado, omitempty"`
}
func (usuario* Usuario)ToJson()(json_payload []byte) {
	json_payload, err := json.Marshal(usuario)
	logging.FailOnError(err, "Error encoding Usuario struct")
	return
}
func (usuario* Usuario)Set(json_payload []byte){
	err := json.Unmarshal(json_payload, usuario)
	logging.FailOnError(err, "Failed to convert json to Usuario")
}

type omit *struct{}
type UsuarioPublico struct {
	*Usuario
	Password omit `json:"password,omitempty"`
}
/*
// when you want to encode your user:
json.Marshal(PublicUser{
User: user,
})
*/
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

/*
func hashSample() {
	userPassword1 := "some user-provided password"

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword1), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))
	// Store this "hash" somewhere, e.g. in your database

	// After a while, the user wants to log in and you need to check the password he entered
	userPassword2 := "some user-provided password"
	hashFromDatabase := hash

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(userPassword2)); err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}

	fmt.Println("Password was correct!")
}
*/