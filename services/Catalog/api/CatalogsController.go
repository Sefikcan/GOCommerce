package controller

import (
	"catalog/infrastructure"
	"catalog/infrastructure/entities"
	"catalog/models"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	products := new([]entities.Product)

	if result := Connection.DB.Find(&products); result.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Product not found!",
		})
	}

	return c.JSON(fiber.Map{
		"data": &products,
	})
}

func GetProductById(c *fiber.Ctx) error {
	product := new(entities.Product)

	if result := Connection.DB.First(&product, c.Params("id")); result.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Product not found!",
		})
	}

	return c.JSON(fiber.Map{
		"data": &product,
	})
}

func CreateProduct(c *fiber.Ctx) error{
	product := new(entities.Product)

	if err := c.BodyParser(product); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	if result := Connection.DB.Create(&product); result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": result.Error,
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"data": &product,
		"message": "Product created successfully",
	})
}

func DeleteProduct(c *fiber.Ctx) error{
	product := new(entities.Product)

	if result := Connection.DB.First(&product, c.Params("id")); result.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Product not found!",
		})
	}

	if response := Connection.DB.Delete(&product, c.Params("id")); response.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
		"message": response.Error,
		})
	}

	c.Status(fiber.StatusNoContent)
	return c.JSON(fiber.Map{})
}

func UpdateProduct(c *fiber.Ctx) error{
	product := new(entities.Product)
	updateProduct := new(models.UpdateProduct)

	if err := c.BodyParser(&updateProduct); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	if result := Connection.DB.First(&product, c.Params("id")); result.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Product not found!",
		})
	}

	product.Price = updateProduct.Price
	product.Name = updateProduct.Name
	product.Quantity = updateProduct.Quantity

	if response := Connection.DB.Save(&product); response.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": response.Error,
		})
	}

	return c.JSON(fiber.Map{
		"data": &product,
		"message": "Product updated successfully",
	})
}
