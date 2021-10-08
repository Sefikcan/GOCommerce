package entities

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