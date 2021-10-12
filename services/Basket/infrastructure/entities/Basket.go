package entities

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

type Basket struct {
	UserId        int     `json:"userId"`
	BasketItem    []BasketItem `json:"basketItem"`
}

type BasketItem struct {
	ProductId     int 	  `json:"productId"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	BasketTotal   float64 `json:"basketTotal"`
}

func(basket *Basket) SetBasketTotal(){
	for i:=0; i< len(basket.BasketItem); i++ {
		basket.BasketItem[i].BasketTotal = basket.BasketItem[i].Price * float64(basket.BasketItem[i].Quantity)
	}
}

func (basket Basket) ValidateBasket() error{
	for i:=0; i< len(basket.BasketItem); i++ {
		return validation.ValidateStruct(&basket.BasketItem[i],
			validation.Field(&basket.BasketItem[i].ProductId, validation.Min(1)),
			validation.Field(&basket.BasketItem[i].Price, validation.Min(float64(1))),
			validation.Field(&basket.BasketItem[i].Quantity, validation.Min(12)))
	}

	return validation.ValidateStruct(&basket,
		validation.Field(&basket.UserId, validation.Required),
		validation.Field(&basket.BasketItem, validation.NotNil))
}