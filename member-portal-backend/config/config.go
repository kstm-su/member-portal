package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	File struct {
		Base string `yaml:"base"`
	} `yaml:"file"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`

	JWT struct {
		Key struct {
			PrivateKey string `yaml:"private_key"`
			PublicKey  string `yaml:"public_key"`
		} `yaml:"key"`
		Issuer string `yaml:"issuer"`
		Realm  string `yaml:"realm"`
		KeyId  string `yaml:"key_id"`
	} `yaml:"jwt"`
}

func init() {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.host", "localhost")

	viper.SetDefault("file.base", "/app")

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "password")

	viper.SetDefault("jwt.key.private_key", "private.pem")
	viper.SetDefault("jwt.key.public_key", "public.pem")
	viper.SetDefault("jwt.issuer", "localhost")
	viper.SetDefault("jwt.realm", "localhost")
	viper.SetDefault("jwt.key_id", "key")
}

var Cfg Config

func Load(configFile string) (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err = viper.WriteConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to write default config: %s \n", err)
		}
		slog.Info("設定ファイルが存在しないため、デフォルト設定ファイルを作成しました。")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー: %s \n", err)
	}
	slog.Info("設定ファイルを読み込みました。")

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s \n", err)
	}
	slog.Info("設定ファイルを構造体に変換しました。")

	return &Cfg, nil
}
