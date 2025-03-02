package initializers

import "jwt-authentication-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
