package config

var Server string
var Port string
var Api_root string
var Content_type string

func init(){
	Server = "localhost"
	Port = ":8080"
	Api_root = "/api/"
	Content_type = "application/json"
}