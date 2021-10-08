package constants

const (
	ORDERINGCONNSTRING = "mongodb://localhost:27017/"
	ORDERINGDB = "OrderDB"
	ORDERINGCL = "ORDER"
)

type OrderStatus int

const (
	PENDING = iota
	SHIPPED
	FAILED
	SUCCESS
)
