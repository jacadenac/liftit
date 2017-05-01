package main

import (
	"log"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/jacadenac/liftit/logging"
	"net/http"
	"github.com/jacadenac/liftit/contracts"
)

var conn *amqp.Connection
var ch *amqp.Channel
var replies <-chan amqp.Delivery

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	println("paso1")
	consuming("GET", Get)
	println("paso10")
	consuming("GETBYID", GetById)
	consuming("POST", Post)
	consuming("PUT", Put)
	consuming("DELETE", Put)
}


func consuming(routing_key string, bdtransaction func(request []byte)(response []byte, httpStatus int)) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	logging.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	logging.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	println("paso2")
	q, err := ch.QueueDeclare(
		routing_key, 		// name
		false,       	// durable
		false,       	// delete when usused
		false,       	// exclusive
		false,       	// no-wait
		nil,         	// arguments
	)
	println("paso2.1")
	logging.FailOnError(err, "Failed to declare a queue")
	println("paso3")
	err = ch.Qos(
		1,     	// prefetch count
		0,     	// prefetch size
		false, 		// global
	)
	println("paso3.1")
	logging.FailOnError(err, "Failed to set QoS")
	println("paso4")
	msgs, err := ch.Consume(
		q.Name, 		// queue
		"",     	// consumer
		false,  	// auto-ack
		false,  	// exclusive
		false,  	// no-local
		false,  		// no-wait
		nil,    		// args
	)
	println("paso4.1")
	logging.FailOnError(err, "Failed to register a consumer")
	println("paso5")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//n, err := strconv.Atoi(string(d.Body))
			//logging.FailOnError(err, "Failed to convert body to integer")

			//log.Printf(" [.] fib(%d)", n)
			response, _ := bdtransaction(d.Body)

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			logging.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

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
		response, err = json.Marshal(usuario)
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
	response = []byte(`{"Post":"usuario registrado"}`)
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
