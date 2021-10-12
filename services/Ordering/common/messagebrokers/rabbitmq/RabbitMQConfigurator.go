package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"ordering/common/messagebrokers/rabbitmq/models"
)

func PrepareRabbitMQ(model models.AmqpModel) error {
	//Create rabbitmq connection
	connectRabbitMQ, err := amqp.Dial(model.RabbitMQUrl)
	if err != nil {
		fmt.Println("test")
		panic(err)
	}

	//Open rabbitmq channel
	channel, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}

	//We define the queue where we can subscribe to For Producers and consumers.
	_, err = channel.QueueDeclare(
		model.QueueName,
		model.Durable,
		model.AutoDelete,
		model.Exclusive,
		model.NoWait,
		nil,
	)
	if err != nil {
		panic(err)
	}

	message := amqp.Publishing{
		ContentType: model.ContentType,
		Body:        model.Body,
	}

	pErr := channel.Publish(
		model.ExchangeName,
		model.QueueName,
		model.Mandatory,
		model.Immediate,
		message)

	return pErr
}

