package Consumers

import (
	"email/common/constants"
	"email/models"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func SendEmail(){
	//TODO: rabbitmq consume işlemlerini ortaklaştır
	connectRabbitMQ, err := amqp.Dial(constants.RABBITMQURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	emailChannel, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer emailChannel.Close()

	messages, err := emailChannel.Consume(
		constants.CREATEORDERQUEUENAME,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		panic(err)
	}

	//Send Email operation blabla

	var order models.Order

	for msg := range messages {
		err := json.Unmarshal(msg.Body, &order)
		if err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Println("Order Customer Id:", order.UserId)
	}
}
