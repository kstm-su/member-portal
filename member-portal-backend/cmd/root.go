package cmd

import (
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/kstm-su/Member-Portal/backend/crypto"
	"github.com/kstm-su/Member-Portal/backend/database"
	"github.com/kstm-su/Member-Portal/backend/router"
	"github.com/spf13/cobra"
)

var configFile string

const name = "member-portal"

var rootCmd = &cobra.Command{
	Use:   name,
	Short: "Backend server for the OAuth2 server",
}

func init() {
	// コマンドフラグの設定
	flags := rootCmd.PersistentFlags()
	// 設定ファイルのパスを指定するフラグ --config, -c
	flags.StringVarP(&configFile, "config", "c", "/app/data/config.yaml", "config file path (default is /app/data/config.yaml)")
}

func Execute() error {
	// コマンドの実行
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	c, err := config.Load(configFile)
	if err != nil {
		print(err.Error())
		return err
	}
	//キーペアの生成
	crypto.Init(*c)
	// データベースの初期化
	database.InitDatabase(c)
	// ルーターの実行
	router.Execute(c)
	return nil
}
