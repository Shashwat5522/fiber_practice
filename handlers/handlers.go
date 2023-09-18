package handlers

import (
	"database/sql"
	"fiber/database"
	"fiber/model"
	"log"

	"github.com/gofiber/fiber"
)

func GetAllProducts(c *fiber.Ctx) {
	rows, err := database.DB.Query("select name,description,category,amount FROM products order by name")
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
	defer rows.Close()

	result := model.Products{}
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
			return
		}
		result.Products = append(result.Products, product)

	}
	if err := c.JSON(&fiber.Map{
		"success": true,
		"product": result,
		"message": "All products returned successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
}

func GetSingleProduct(c *fiber.Ctx) {
	id := c.Params("id")
	product := model.Product{}

	row, err := database.DB.Query("select * from products where id=$1", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	defer row.Close()

	for row.Next() {
		switch err := row.Scan(&id, &product.Amount, &product.Name, &product.Description, &product.Category); err {
		case sql.ErrNoRows:
			log.Println("No rows were returned")
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		case nil:
			log.Println(product.Name, product.Description, product.Category, product.Amount)

		default:
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}
	if err := c.JSON(&fiber.Map{
		"success": false,
		"message": "Sucessfully fetched product",
		"product": product,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
}

func CreateProduct(c *fiber.Ctx) {
	p := new(model.Product)

	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	res, err := database.DB.Query("Insert INTO products(name,description,category,amount) VALUES ($1,$2,$3,$4)", p.Name, p.Description, p.Category, p.Amount)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	log.Println(res)

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "product successfully create",
		"product": p,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return
	}
}

func DeleteProduct(c *fiber.Ctx) {
	id := c.Params("id")
	res, err := database.DB.Query("DELETE FROM products where id=$1", id)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
	log.Println(res)

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "product deleted successfully",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"Success": false,
			"error":   err,
		})
		return
	}
}

func UpdateProduct(c *fiber.Ctx) {
	id := c.Params("id")
	var product model.Product

	if err := c.BodyParser(&product); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	log.Println(product)
	res, err := database.DB.Query("UPDATE products SET name=$1,description=$2,category=$3,amount=$4 WHERE id=$5", product.Name, product.Description, product.Category, product.Amount, id)

	if err != nil {
		log.Println(err)
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	log.Println(res)

	if err:=c.JSON(&fiber.Map{
		"success":true,
		"message":"Product updated successfully!!!",
	});err!=nil{
		c.JSON(&fiber.Map{
			"success":false,
			"message":err,
		})
		return
	}

}
