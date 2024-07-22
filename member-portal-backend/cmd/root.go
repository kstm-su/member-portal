package cmd

import (
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/kstm-su/Member-Portal/backend/router"
	"github.com/spf13/cobra"
)

var configFile string

const name = "member-portal"
const file = "." + name + ".yaml"

var rootCmd = &cobra.Command{
	Use:   name,
	Short: "Backend server for the OAuth2 server",
}

func init() {
	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&configFile, "config", "c", "/app/config.yaml", "config file path (default is /app/config.yaml)")
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	c, err := config.Load(configFile)
	if err != nil {
		print(err.Error())
		return err
	}
	router.Execute(c)
	return nil
}
