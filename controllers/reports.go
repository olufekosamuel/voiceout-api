package controllers

import (
	"encoding/json"

	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/olufekosamuel/voiceout-api/helpers"
	"github.com/olufekosamuel/voiceout-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterAdmin(c *fiber.Ctx) error {
	c.Accepts("application/json")

	admin := models.NewAdmin()

	if err := json.Unmarshal(c.Body(), &admin); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if len(admin.Email) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": true, "msg": "Email can not be empty"})
	}

	if len(admin.Password) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": true, "msg": "Password must be provided"})
	}

	if len(admin.FirstName) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": true, "msg": "Firstname must be provided"})
	}

	if len(admin.Surname) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": true, "msg": "Surname must be provided"})
	}

	if len(admin.PhoneNumber) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": true, "msg": "Phone number must be provided"})
	}

	// check if user exists

	collection := mgm.Coll(&models.Admin{})

	var admins []models.Admin

	err := collection.SimpleFind(&admins, bson.D{{"email", admin.Email}})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if len(admins) != 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"msg":   "Admin with that email already exists",
		})
	}

	admin.Password, _ = helpers.HashPassword(admin.Password)

	err = mgm.Coll(admin).Create(admin)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	token, err := helpers.GenerateToken(admin.Email)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	admin.Password = ""

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"admin": admin,
			"token": token,
		},
	})

}

func LoginAdmin(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var admin models.Admin

	if err := json.Unmarshal(c.Body(), &admin); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if len(admin.Email) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": true, "msg": "Email can not be empty"})
	}

	if len(admin.Password) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": true, "msg": "Password must be provided"})
	}

	collection := mgm.Coll(&models.Admin{})

	var admins []models.Admin

	err := collection.SimpleFind(&admins, bson.D{{"email", admin.Email}})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if len(admins) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"msg":   "User with that email does not exist",
		})
	}

	if !helpers.CheckPasswordHash(admin.Password, admins[0].Password) {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   "Incorrect email or password",
		})
	}

	token, err := helpers.GenerateToken(admin.Email)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	admins[0].Password = ""

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"admin": admins[0],
			"token": token,
		},
	})
}

// func UpdateAdminPassword(c *fiber.Ctx) error {
// 	c.Accepts("application/json")

// 	id := c.Params("id")

// 	admin := &models.Admin{}

// 	collection := mgm.Coll(admin)

// 	err := collection.FindByID(id, admin)

// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	var usr models.NewPassword

// 	if err := json.Unmarshal(c.Body(), &usr); err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	if !helpers.CheckPasswordHash(usr.OldPassword, admin.Password) {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "Your current password is incorrect",
// 		})
// 	}

// 	// write update logic here
// 	admin.Password, _ = helpers.HashPassword(usr.Password)

// 	err = collection.Update(admin)

// 	if err := json.Unmarshal(c.Body(), &usr); err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"error": false,
// 		"data":  admin,
// 		"msg":   "Admin password updated successfully",
// 	})
// }

func CreateReport(c *fiber.Ctx) error {
	c.Accepts("application/json")

	newreport := models.NewReport()

	newreport.Name = c.FormValue("name")
	newreport.Department = c.FormValue("department")
	newreport.Description = c.FormValue("description")
	newreport.Anonymous = c.FormValue("anonymous")
	newreport.Fullname = c.FormValue("fullname")
	newreport.Email = c.FormValue("email")
	newreport.Phone = c.FormValue("phone")

	file, _ := c.FormFile("file")
	file_url, err := helpers.UploadFile(file, "candidate")
	newreport.File = file_url

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// store user details to database

	err = mgm.Coll(newreport).Create(newreport)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  newreport,
	})
}

func CreateDepartment(c *fiber.Ctx) error {
	c.Accepts("application/json")

	department := models.NewDepartment()

	if err := json.Unmarshal(c.Body(), &department); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err := mgm.Coll(department).Create(department)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  department,
	})
}

func GetDepartment(c *fiber.Ctx) error {
	var department []models.Department

	collection := mgm.Coll(&models.Department{})

	err := collection.SimpleFind(&department, bson.D{})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  department,
	})

}
