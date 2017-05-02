package main

import (
	"sync"
	"database/sql"
	"github.com/jacadenac/liftit/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jacadenac/liftit/contracts"
	"fmt"
	"encoding/json"
)

//Se implementa estructura usuarioServ bajo patron singleton
type database struct{
	 Conn *sql.DB
}
var instance *database
var once sync.Once
func GetDatabase() *database {
	once.Do(func() {
		// Create the database handle, confirm driver is present
		db, err := sql.Open("mysql", "root:passbd@tcp(localhost:3307)/sample?charset=utf8")
		logging.FailOnError(err, "Error conectando con la base de datos")

		// Test the connection to the database
		err = db.Ping()
		logging.FailOnError(err, "Error al hacer ping...")
		instance = &database{
			db,
		}

		// Initialize Table and Raw
		//Init()
	})
	return instance
}
func GetConn() *sql.DB{
	data := GetDatabase()
	return data.Conn
}

func Init() {
	// sql.DB should be long lived "defer" closes it once this function ends
	//defer db.Close()
	db := GetConn()
	query := `CREATE OR REPLACE TABLE usuario (
		id INT(11) NOT NULL AUTO_INCREMENT,
		nombre CHAR(100) NOT NULL,
		email CHAR(100) NOT NULL,
		password CHAR(60) NOT NULL,
		verificado TINYINT(1) NULL DEFAULT NULL,
		telefono INT(11) NOT NULL,
		pais CHAR(100) NULL DEFAULT NULL,
		ciudad CHAR(100) NULL DEFAULT NULL,
		direccion CHAR(200) NULL DEFAULT NULL,
		fecha_creado TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	)
	COLLATE='latin1_swedish_ci'
	ENGINE=InnoDB
	;`
	stmtIns, err := db.Prepare(query)
	//stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
	logging.FailOnError(err, "Error preparando sentencia para crear tabla sql")
	defer stmtIns.Close()
	_, err = stmtIns.Exec()
	if err != nil {
		if(err.Error() != "Error 1050: Table 'usuario' already exists"){
			logging.FailOnError(err, "Error ejecutando sentencia para crear tabla sql")
		}
	}

	query = `INSERT INTO Usuario
		(nombre, email, password, verificado, telefono, pais, ciudad, direccion, fecha_creado)
		VALUES(
			'Juan David Alejandro Cadena Cedano',
			'jacadenac@unal.edu.co',
			'juan123',
			true,
			301659808,
			'Colombia',
			'Bogotá',
			'Calle falsa 123',
			'2016-05-01 23:24:00'
	)`
	stmtIns, err = db.Prepare(query)
	logging.FailOnError(err, "Error preparando sentencia para crear usuario sql")
	_, err = stmtIns.Exec()
	logging.FailOnError(err, "Error ejecutando sentencia para crear usuario sql")

	/*
	// Prepare statement for reading data


	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)
	*/
}

func main(){
	_ = GetConn()
	Init()
	//usuario, ok, err := GetById(1)
	//fmt.Println(usuario, ok, err)
}


func Get()(ok bool, err error){
	return
}

func GetById(id int)(usuario contracts.Usuario, ok bool, err error){
	db := GetConn()
	query := "SELECT nombre, email, password, verificado, telefono, pais, ciudad, direccion, fecha_creado FROM Usuario WHERE id = ?"
	row := db.QueryRow(query, id)
	//rows, _ := db.Query(query, id)

	err = row.Scan(
		&usuario.Nombre,
		&usuario.Email,
		&usuario.Password,
		&usuario.Verificado,
		&usuario.Telefono,
		&usuario.Pais,
		&usuario.Ciudad,
		&usuario.Direccion,
		&usuario.FechaCreado,
	)
	logging.FailOnError(err, "Error escaneando filas")


	/*
	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)
	*/
	return
}

func Save(usuario contracts.Usuario)(ok bool, err error){
	return
}

func Update(id int, usuario contracts.Usuario)(ok bool, err error) {
	return
}

func Delete(id int)(ok bool, err error) {
	return
}


func GetJSON(sqlString string) ([]byte, error) {
	db := GetConn()
	rows, err := db.Query(sqlString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonData))
	return jsonData, nil
}