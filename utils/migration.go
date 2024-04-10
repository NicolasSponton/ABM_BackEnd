package utils

import (
	"ABM_Clientes/database"
	"ABM_Clientes/models"
)

func MigrateSchemas() {
	db := database.GetConnection()
	db.AutoMigrate(&models.Cliente{})

}
