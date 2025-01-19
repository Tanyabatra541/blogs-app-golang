package controller

import (
	"log"

	"github.com/Tanyabatra541/blogs-app-golang/database"
	"github.com/Tanyabatra541/blogs-app-golang/model"
	"github.com/gofiber/fiber/v2"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "List of all blogs",
	}
	db := database.DBConn
	var records []model.Blog
	db.Find(&records)
	context["blog_records"] = records
	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "Add Blog",
	}
	record := new(model.Blog)
	if err := c.BodyParser(&record); err != nil{
		log.Println("Error parsing the request body.")
		context["statusText"] = ""
		context["msg"] = "Error parsing the request body."
	}
	result := database.DBConn.Create(record)
	if result.Error != nil{
		log.Println("Error saving the record.")
	}
	context["msg"] = "Blog added successfully."
	context["data"] = record
	c.Status(201)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "Update Blog",
	}
	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"msg": "Delete Blog for the given ID",
	}
	c.Status(200)
	return c.JSON(context)
}