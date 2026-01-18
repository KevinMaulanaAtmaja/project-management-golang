package seed

import (
	"log"

	"github.com/KevinMaulanaAtmaja/project-management-golang/config"
	"github.com/KevinMaulanaAtmaja/project-management-golang/models"
	"github.com/KevinMaulanaAtmaja/project-management-golang/utils"
	"github.com/google/uuid"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name:     "Super admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
		PublicID: uuid.New(),
	}

	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed too seed admin", err)
	} else {
		log.Println("Admin user seeded")
	}
}
