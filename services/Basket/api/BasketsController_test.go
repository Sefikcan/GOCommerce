package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

var c *fiber.Ctx

func TestGetBasketByUserId(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "http:localhost:8002/api/v1/baskets/1234",
			expectedCode: 200,
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Create route with GET method for test
	app.Get("http:localhost:8002/api/v1/baskets/1234", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

/*func TestAddOrUpdateBasket(t *testing.T) {
	app := fiber.New()
	app.Post("/api/v1/baskets/", AddOrUpdateBasket)

	body := "{\n\t\"userId\":1234,\n\t\"basketItem\":[\n\t{\n\t\t\"price\": 123,\n\t\t\"quantity\":13\n\t},\n\t{\n\t\t\"productId\":5,\n\t\t\"price\": 122,\n\t\t\"quantity\":1\n\t}\n\t]\n}"
	req:=httptest.NewRequest("POST", "http://localhost:8002/api/v1/baskets/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Method = "POST"

	resp, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}*/

/*func TestAddOrUpdateBasket(t *testing.T) {
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

/*func TestRemoveBasketByUserId(t *testing.T) {
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
}*/