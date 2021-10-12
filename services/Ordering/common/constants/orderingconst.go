package constants

const (
	ORDERINGCONNSTRING = "mongodb://localhost:27017/"
	ORDERINGDB = "OrderDB"
	ORDERINGCL = "ORDER"
	RABBITMQURL = "amqp://guest:guest@localhost:5672/"
	CREATEORDERQUEUENAME = "CreateOrderQueueName"
)

type OrderStatus int

const (
	PENDING = iota
	SHIPPED
	FAILED
	SUCCESS
)
