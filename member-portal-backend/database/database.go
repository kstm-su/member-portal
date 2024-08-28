package database

import (
	"fmt"
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/kstm-su/Member-Portal/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase(c *config.Config) {
	var err error
	log.Println("Connecting to database")
	switch c.Database.Type {
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(c.Database.SQLite.Path), &gorm.Config{})
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d", c.Database.Postgres.Host, c.Database.Postgres.User, c.Database.Postgres.Password, c.Database.Postgres.Port)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		log.Fatal("failed to connect database")
	}
	log.Println("Connected to database")

	// Auto-migrate the models
	err = DB.AutoMigrate(&models.User{}, &models.Auth{}, &models.Role{}, &models.Affiliation{}, &models.Faculty{}, &models.Contact{}, &models.Name{}, &models.Profile{}, &models.ActivityLog{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
}
