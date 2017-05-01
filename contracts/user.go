package contracts

import (
	"time"
	"strconv"
	//"golang.org/x/crypto/bcrypt"
)

var UsuarioStore = make(map[string]Usuario)
var ID int

type Usuario struct {
	Nombre		string		`json:"nombre" binding:"required"`
	Email		string		`json:"description" binding:"required"`
	Password 	string		`json:"password" binding:"required"`
	Verificado	bool		`json:"verificado,omitempty"`
	Telefono	int		`json:"telefono" binding:"required"`
	Pais		string		`json:"pais",omitempty`
	Ciudad		string		`json:"Ciudad",omitempty`
	Direccion	string		`json:"Direccion",omitempty`
	CreatedAt 	time.Time 	`json:"created_at,omitempty"`
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