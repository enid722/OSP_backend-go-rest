package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm/v3"
	"github.com/enid722/OSP_backend-go-rest/controllers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	app.Get("/api/surveys", controllers.GetAllSurveys)
	app.Get("/api/surveys/:id", controllers.GetSurveyByID)
	app.Post("/api/surveys", controllers.CreateSurvey)
	app.Delete("/api/surveys/:id", controllers.DeleteSurvey)

	log.Fatal(app.Listen(":3000"))
}

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}

	err := mgm.SetDefaultConfig(nil, "OSP", options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
}
