package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	Connection "ordering/infrastructure"
	"ordering/infrastructure/entities"
	"strconv"
)

//bson.D -> Document
//bson.A -> Array
//bson.M -> Map

func GetOrderByUserId(c *fiber.Ctx) error {
	ordering := &entities.Order{}
	userId,convErr := strconv.Atoi(c.Params("userId"))
	if convErr != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": convErr.Error(),
		})
	}

	err := Connection.MongoCollection.FindOne(c.Context() ,bson.M{"user_id": userId}).Decode(ordering)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Order is empty!",
		})
	}

	return c.JSON(fiber.Map{
		"data": &ordering,
	})
}

func CreateOrder(c *fiber.Ctx) error{
	ordering := new(entities.Order)

	json.Unmarshal([]byte(c.Body()), &ordering)

	ordering.SetOrderTotal()
	ordering.OrderStatus = 1

	resp,err := Connection.MongoCollection.InsertOne(c.Context() , ordering)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ordering.Id = resp.InsertedID.(primitive.ObjectID).String()
	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"data": &ordering,
		"message": "Order created successfully",
	})
}

func DeleteOrder(c *fiber.Ctx) error{
	orderId, _ := primitive.ObjectIDFromHex(c.Params("id"))

	_, err := Connection.MongoCollection.DeleteOne(c.Context() , bson.M{"_id": orderId})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Status(fiber.StatusNoContent)
	return c.JSON(fiber.Map{})
}

//TODO: Remove Order Item

func UpdateOrder(c *fiber.Ctx) error {
	ordering := new(entities.Order)
	orderId, _ := primitive.ObjectIDFromHex(c.Params("id"))
	json.Unmarshal([]byte(c.Body()), &ordering)

	ordering.SetOrderTotal()

	update := bson.M{
		"$set": ordering,
	}

	_, err := Connection.MongoCollection.UpdateOne(c.Context() ,bson.M{"_id": orderId}, update)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": &ordering,
		"message": "Order updated successfully",
	})
}