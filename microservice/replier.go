package main

import (
	"log"
	"github.com/streadway/amqp"
	"github.com/jacadenac/liftit/logging"
	"flag"
	"github.com/jacadenac/liftit/config"
	"github.com/jacadenac/liftit/microservice/controler"
)

var conn *amqp.Connection
var ch *amqp.Channel

func main() {
	var err error
	flag.Parse()
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	logging.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err = conn.Channel()
	logging.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs1 := createQueueAndConsumer("GET")
	msgs2 := createQueueAndConsumer("GETBYID")
	msgs3 := createQueueAndConsumer("POST")
	msgs4 := createQueueAndConsumer("PUT")
	msgs5 := createQueueAndConsumer("DELETE")

	forever := make(chan bool)

	go responder(msgs1, controler.Get)
	go responder(msgs2, controler.GetById)
	go responder(msgs3, controler.Post)
	go responder(msgs4, controler.Put)
	go responder(msgs5, controler.Delete)

	log.Println(" [*] Awaiting RPC requests")
	<-forever
}

func createQueueAndConsumer(queue_name string)(msgs <-chan amqp.Delivery){
	queue, err := ch.QueueDeclare(
		queue_name,		// name
		false,       	// durable
		false,       	// delete when usused
		false,       	// exclusive
		false,       	// no-wait
		nil,         	// arguments
	)
	logging.FailOnError(err, "Failed to declare a queue")
	err = ch.Qos(
		1,     	// prefetch count
		0,     	// prefetch size
		false, 		// global
	)
	logging.FailOnError(err, "Failed to set QoS")
	//log.Println("Microservice: creó queue = ", queue.Name)

	msgs, err = ch.Consume(
		queue.Name, 		// queue
		"",     	// consumer
		false,  	// auto-ack
		false,  	// exclusive
		false,  	// no-local
		false,  		// no-wait
		nil,    		// args
	)
	logging.FailOnError(err, "Failed to register a consumer")
	//log.Println("Creó consumidor de peticiones = ", queue.Name)
	return
}

func responder(msgs <-chan amqp.Delivery, bdtransaction func(request []byte)(response []byte, httpStatus int)){
	for d := range msgs {
		//log.Println("¡Llegó un mensaje!")
		//log.Println("d.ReplyTo = ", d.ReplyTo)
		//log.Println("d.CorrelationId = ", d.CorrelationId)

		response, _ := bdtransaction(d.Body)
		err := ch.Publish(
			"",        	// exchange
			d.ReplyTo, 		// routing key
			false,     	// mandatory
			false,     	// immediate
			amqp.Publishing{
				ContentType:   config.Content_type,
				CorrelationId: d.CorrelationId,
				Body:          response,
			})
		logging.FailOnError(err, "Failed to publish a message")
		d.Ack(false)
	}
}

