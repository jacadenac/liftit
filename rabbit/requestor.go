package rabbit

import (
	"time"
	"math/rand"
	"github.com/streadway/amqp"
	"github.com/jacadenac/liftit/logging"
	"github.com/jacadenac/liftit/config"
)

//var Conn *amqp.Connection
//var Ch *amqp.Channel
//var queue amqp.Queue
//var messages <-chan amqp.Delivery
//var err error

func Publish(routing_key string, request []byte) (response []byte, err error){

	conn, err := amqp.Dial(*config.AmqpURI)
	logging.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()


	ch, err := conn.Channel()
	logging.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	logging.FailOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	logging.FailOnError(err, "Failed to register a consumer")

	corrId := randomString(32)
	err = ch.Publish(
		"",          	// exchange
		routing_key, 		// routing key
		false,       	// mandatory
		false,       	// immediate
		amqp.Publishing{
			ContentType:   config.Content_type,
			CorrelationId: corrId,
			ReplyTo:       queue.Name,
			Body:          []byte(request),
		})
	logging.FailOnError(err, "Failed to publish a message")

	for d := range messages {
		if corrId == d.CorrelationId {
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
