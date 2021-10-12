package entities

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func (p Product) ValidateProduct() error{
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Length(3,50)),
		validation.Field(&p.Price, validation.Min(0)),
		validation.Field(&p.Quantity, validation.Min(1)))
}