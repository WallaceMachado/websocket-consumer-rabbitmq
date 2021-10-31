package queue

import (
	"github.com/streadway/amqp"
	"github.com/wallacemachado/websocket-consumer-rabbitmq/src/config"
)

func Connect() *amqp.Channel {
	dsn := "amqp://" + config.RabbitmqUser + ":" + config.RabbitmqPass + "@" + config.RabbitmqHost + ":" + config.RabbitmqPort + config.RabbitmqVhost

	conn, err := amqp.Dial(dsn)
	if err != nil {
		panic(err.Error())
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}
	return channel
}

func StartConsuming(queue string, ch *amqp.Channel, in chan []byte) {

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err.Error())
	}

	msgs, err := ch.Consume(
		q.Name,
		"websocket",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		panic(err.Error())
	}

	go func() {
		for m := range msgs {
			in <- []byte(m.Body)
		}
		close(in)
	}()
}
