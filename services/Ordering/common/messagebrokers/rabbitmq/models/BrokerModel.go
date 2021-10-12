package models

type AmqpModel struct {
	RabbitMQUrl 	string
	QueueName 		string
	ExchangeName 	string
	Durable 		bool
	AutoDelete 		bool
	Exclusive 		bool
	NoWait    		bool
	Mandatory 		bool
	Immediate 		bool
	ContentType 	string
	Body 			[]byte
}
