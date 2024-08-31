package database

import (
	"fmt"
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/kstm-su/Member-Portal/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
)

var DB *gorm.DB

func InitDatabase(c *config.Config) {
	var err error
	slog.Info("データベースへの接続を開始します")
	switch c.Database.Type {
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(c.Database.SQLite.Path), &gorm.Config{})
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d", c.Database.Postgres.Host, c.Database.Postgres.User, c.Database.Postgres.Password, c.Database.Postgres.Port)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		slog.Error("データベースへの接続に失敗しました")
	}
	slog.Info("データベースへの接続が完了しました")

	// Auto-migrate the models
	err = DB.AutoMigrate(&models.Users{}, &models.Auth{}, &models.Role{}, &models.Affiliation{}, &models.Faculty{}, &models.Contact{}, &models.Name{}, &models.Profile{}, &models.ActivityLog{})
	if err != nil {
		slog.Error("データベースのマイグレーションに失敗しました")
	}
}
