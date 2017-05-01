package requestor

import (
	"time"
	"math/rand"
	"github.com/streadway/amqp"
	"github.com/jacadenac/liftit/logging"
	"github.com/jacadenac/liftit/config"
)

var Conn *amqp.Connection
var Ch *amqp.Channel
var messages <-chan amqp.Delivery

func Publish(routing_key string, request []byte) (response []byte, err error){
	Conn, err = amqp.Dial(*config.AmqpURI)
	logging.FailOnError(err, "Failed to connect to RabbitMQ")
	defer Conn.Close()

	Ch, err = Conn.Channel()
	logging.FailOnError(err, "Failed to open a channel")
	defer Ch.Close()

	cola_respuestas, err := Ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	//log.Println("Requestor: creó una cola con nombre = ", cola_respuestas.Name)
	logging.FailOnError(err, "Failed to declare a cola_respuestas")

	messages, err = Ch.Consume(
		cola_respuestas.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	logging.FailOnError(err, "Failed to register a consumer")
	//log.Println("Creó consumidor de cola_respuestas = ", cola_respuestas.Name)

	corrId := randomString(32)
	err = Ch.Publish(
		"",          	// exchange
		routing_key, 		// routing key
		false,       	// mandatory
		false,       	// immediate
		amqp.Publishing{
			ContentType:   config.Content_type,
			CorrelationId: corrId,
			ReplyTo:       cola_respuestas.Name,
			Body:          []byte(request),
		})
	logging.FailOnError(err, "Failed to publish a message")
	//log.Println("¡Envió mensaje!")
	//log.Println("d.ReplyTo = ", cola_respuestas.Name)
	//log.Println("d.CorrelationId = ", corrId)

	for d := range messages {
		if corrId == d.CorrelationId {
			//log.Printf("¡Llegó una respuesta!")
			//log.Println("d.CorrelationId = ", d.CorrelationId)
			response = []byte(d.Body)
			break
		}
	}
	return
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return time.Now().String()+"-"+string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
