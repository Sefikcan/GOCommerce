package models

//TODO: Ortak modeller ortak yere taşınacak
type Order struct {
	Id           string      		 	 `json:"_id"`
	OrderTotal   float64    		     `json:"orderTotal"`
	UserId       int		  		     `json:"userId"`
	OrderItem    []OrderItem 	         `json:"orderItem"`
}

type OrderItem struct {
	ProductId int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
