package controller

import (
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