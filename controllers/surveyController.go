package controllers

import (
	"crypto/rand"
	"fmt"

	"github.com/Kamva/mgm/v3"
	"github.com/enid722/OSP_backend-go-rest/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllSurveys - GET /api/surveys
func GetAllSurveys(c *fiber.Ctx) error {
	collection := mgm.Coll(&models.Survey{})
	surveys := []models.Survey{}

	err := collection.SimpleFind(&surveys, bson.D{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"surveys": surveys,
	})
}

// GetSurveyByID - GET /api/surveys/:id
func GetSurveyByID(c *fiber.Ctx) error {
	id := c.Params("id")

	survey := &models.Survey{}
	collection := mgm.Coll(survey)

	err := collection.FindByID(id, survey)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Survey not found.",
		})
	}

	return c.JSON(fiber.Map{
		"ok":     true,
		"survey": survey,
	})

}

// CreateSurvey - POST /api/surveys
func CreateSurvey(c *fiber.Ctx) error {
	params := new(struct {
		Title     string
		Questions []models.Question
	})

	c.BodyParser(&params)

	if len(params.Title) == 0 || len(params.Questions) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title / Questions not specified.",
		})

	}

	//generate a random token with length 8
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	token := fmt.Sprintf("%X", b) //convert to string

	survey := models.CreateSurvey(params.Title, token, params.Questions)
	err := mgm.Coll(survey).Create(survey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})

	}

	return c.JSON(fiber.Map{
		"ok":     true,
		"survey": survey,
	})
}

/*
// DeleteSurvey - DELETE /api/surveys/:id
func DeleteSurvey(c *fiber.Ctx) error {
	id := c.Params("id")

	survey := &models.Survey{}
	collection := mgm.Coll(survey)

	err := collection.FindByID(id, survey)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Survey not found.",
		})

	}

	err = collection.Delete(survey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})

	}

	return c.JSON(fiber.Map{
		"ok":   true,
		"survey": survey,
	})
}*/

// DeleteSurvey - DELETE /api/surveys/:id
func DeleteSurvey(c *fiber.Ctx) error {
	id := c.Params("id")

	survey := &models.Survey{}
	collection := mgm.Coll(survey)

	err := collection.FindByID(id, survey)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Survey not found.",
		})

	}

	//Soft delete the survey
	survey.IsDeleted = !survey.IsDeleted

	err = collection.Update(survey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})

	}

	return c.JSON(fiber.Map{
		"ok":     true,
		"survey": survey,
	})
}
