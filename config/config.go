package config

var Server string
var Port string
var Api_root string
var Content_type string

func init(){
	Server = "localhost"
	Port = ":9000"
	Api_root = "/"
	Content_type = "application/json"
}