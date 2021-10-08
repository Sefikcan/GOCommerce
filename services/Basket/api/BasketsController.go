package controller

import (
	"basket/common/constants"
	Connection "basket/infrastructure"
	"basket/infrastructure/entities"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func GetBasketByUserId(c *fiber.Ctx) error {
	basket := new(entities.Basket)
	result, err := Connection.RedisClient.Get(c.Params("userId")).Result()
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Basket is empty!",
		})
	}

	if err := json.Unmarshal([]byte(result), &basket); err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(fiber.Map{
		"data": &basket,
	})
}

func AddOrUpdateBasket(c *fiber.Ctx) error{
	basket := new(entities.Basket)

	if err := c.BodyParser(basket); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	basket.SetBasketTotal()

	p, err := json.Marshal(&basket)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	if result := Connection.RedisClient.Set(strconv.Itoa(basket.UserId), p, time.Minute * constants.CACHEEXPIRATION); result.Err() != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": result.Err().Error(),
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"data": &basket,
	})
}

func RemoveBasketByUserId(c *fiber.Ctx) error{
	basket := new(entities.Basket)

	result, err := Connection.RedisClient.Get(c.Params("userId")).Result()
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Basket is empty!",
		})
	}

	if err := json.Unmarshal([]byte(result), &basket); err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	if  _,err := Connection.RedisClient.Del(c.Params("userId")).Result(); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	c.Status(fiber.StatusNoContent)
	return c.JSON(fiber.Map{})
}

//TODO: Remove BasketItem