package config

import (
	"flag"
)

var Server Host
var Client Host
var Api_root string
var Content_type string
var Access_control AccessControl
var AmqpURI *string

type AccessControl struct{
	Key string
	Value string
}

type Host struct {
	Ip string
	Port string
}

func (host Host) Get() string{
	return host.Ip+host.Port
}

func init(){
	Server = Host{"localhost",":9010"}
	Client = Host{"localhost",":9000"}
	Api_root = "/"
	Content_type = "application/json"
	Access_control = AccessControl{"Access-Control-Allow-Origin", "*"}
	AmqpURI = flag.String("amqp", "amqp://guest:guest@localhost:5672/", "AMQP URI")
}