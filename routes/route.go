package routes

import (
	"log"

	"github.com/KevinMaulanaAtmaja/project-management-golang/config"
	"github.com/KevinMaulanaAtmaja/project-management-golang/controllers"
	"github.com/KevinMaulanaAtmaja/project-management-golang/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App, uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	app.Post("/api/v1/auth/register", uc.Register)
	app.Post("/api/v1/auth/login", uc.Login)

	// JWT protected routes
	api := app.Group("/api/v1", jwtware.New(jwtware.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.Unauthorized(c, "Error Unauthorized", err.Error())
		},
	}))

	useGroup := api.Group("/users")
	useGroup.Get("/page", uc.GetUserPagination) //  /api/v1/users/page
	useGroup.Get("/:id", uc.GetUser)            //  /api/v1/users/:id
	useGroup.Put("/:id", uc.UpdateUser)         //  /api/v1/users/:id
	useGroup.Delete("/:id", uc.DeleteUser)      //  /api/v1/users/:id

}
