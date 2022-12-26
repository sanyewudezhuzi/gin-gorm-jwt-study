package initializers

import "github.com/NotAPigInTheTrefoilHouse/gin-gorm-jwt-study/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
