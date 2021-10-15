package controller

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

var c *fiber.Ctx

func TestGetBasketByUserId(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{ "GetBasketByUserId", args{c}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetBasketByUserId(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBasketByUserId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddOrUpdateBasket(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{ "",args{c: c}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddOrUpdateBasket(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AddOrUpdateBasket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveBasketByUserId(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{ "",args{c: c}, false},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveBasketByUserId(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("RemoveBasketByUserId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}