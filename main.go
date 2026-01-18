package main

import (
	"log"

	"github.com/KevinMaulanaAtmaja/project-management-golang/config"
	"github.com/KevinMaulanaAtmaja/project-management-golang/controllers"
	"github.com/KevinMaulanaAtmaja/project-management-golang/database/seed"
	"github.com/KevinMaulanaAtmaja/project-management-golang/repositories"
	"github.com/KevinMaulanaAtmaja/project-management-golang/routes"
	"github.com/KevinMaulanaAtmaja/project-management-golang/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)
	log.Fatal(app.Listen(":" + port))
}
