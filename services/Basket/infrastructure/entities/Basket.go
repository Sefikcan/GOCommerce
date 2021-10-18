package entities

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

type Basket struct {
	UserId        int     `json:"userId" form:"userId"`
	BasketItem    []BasketItem `json:"basketItem" form:"basketItem"`
}

type BasketItem struct {
	ProductId     int 	  `json:"productId" form:"productId"`
	Price         float64 `json:"price" form:"price"`
	Quantity      int     `json:"quantity" form:"quantity"`
	BasketTotal   float64 `json:"basketTotal" form:"basketTotal"`
}

func(basket *Basket) SetBasketTotal(){
	for i:=0; i< len(basket.BasketItem); i++ {
		basket.BasketItem[i].BasketTotal = basket.BasketItem[i].Price * float64(basket.BasketItem[i].Quantity)
	}
}

func (basket Basket) ValidateBasket() error{
	for i:=0; i< len(basket.BasketItem); i++ {
		if err := validation.ValidateStruct(&basket.BasketItem[i],
			validation.Field(&basket.BasketItem[i].ProductId, validation.Min(1)),
			validation.Field(&basket.BasketItem[i].Price, validation.Min(float64(1))),
			validation.Field(&basket.BasketItem[i].Quantity, validation.Min(12))); err != nil{
			return err
		}
	}

	return validation.ValidateStruct(&basket,
		validation.Field(&basket.UserId, validation.Required),
		validation.Field(&basket.BasketItem, validation.NotNil))
}