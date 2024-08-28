package config

import (
	"crypto/rand"
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"math/big"
	"os"
)

// 設定ファイルの構造体
type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	File struct {
		Base string `yaml:"base"`
	} `yaml:"file"`

	Database struct {
		Type   string `yaml:"type"`
		SQLite struct {
			Path string `yaml:"path"`
		} `yaml:"sqlite"`
		Postgres struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"postgres"`
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

	Password struct {
		Pepper    string `yaml:"pepper"`
		Algorithm string `yaml:"algorithm"`
	} `yaml:"password"`
}

func init() {
	//デフォルト設定ファイルの設定
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.host", "localhost")

	viper.SetDefault("file.base", "/app")

	viper.SetDefault("database.type", "sqlite")

	viper.SetDefault("database.sqlite.path", "/app/db.sqlite3")

	viper.SetDefault("jwt.key.private_key", "private.pem")
	viper.SetDefault("jwt.key.public_key", "public.pem")
	viper.SetDefault("jwt.issuer", "localhost")
	viper.SetDefault("jwt.realm", "localhost")
	viper.SetDefault("jwt.key_id", "key")

	viper.SetDefault("password.pepper", generateRandomString(30))
	viper.SetDefault("password.algorithm", "argon2")
}

var Cfg Config

func Load(configFile string) (*Config, error) {
	//設定ファイの初期化
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	//設定ファイルの存在チェック　ない場合はデフォルト設定ファイルを作成
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err = viper.WriteConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to write default config: %s \n", err)
		}
		slog.Info("設定ファイルが存在しないため、デフォルト設定ファイルを作成しました。")
	}

	//設定ファイルの読み込み
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("設定ファイル読み込みエラー: %s \n", err)
	}
	slog.Info("設定ファイルを読み込みました。")

	//設定ファイルを構造体に変換
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s \n", err)
	}
	slog.Info("設定ファイルを構造体に変換しました。")

	return &Cfg, nil
}

func generateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}
