package entities

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"ordering/common/constants"
)

type Order struct {
	Id           string      		 	 `json:"_id" bson:"_id,omitempty"`
	OrderTotal   float64    		     `json:"orderTotal" bson:"order_total"`
	UserId       int		  		     `json:"userId" bson:"user_id"`
	OrderStatus  constants.OrderStatus   `json:"orderStatus" bson:"order_status"`
	OrderItem    []OrderItem 	         `json:"orderItem" bson:"order_item"`
}

type OrderItem struct {
	ProductId int     `json:"productId" bson:"product_id"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	Price     float64 `json:"price" bson:"price"`
	//TODO: Parçalı shipping status enum
}

func(order *Order) SetOrderTotal(){
	for i:=0; i< len(order.OrderItem); i++ {
		order.OrderTotal += order.OrderItem[i].Price * float64(order.OrderItem[i].Quantity)
	}
}

func (order Order) ValidateBasket() error{
	return validation.ValidateStruct(&order,
		validation.Field(&order.UserId, validation.Required),
		validation.Field(&order.OrderItem, validation.NotNil))
}
