package main

import (
	"github.com/kstm-su/Member-Portal/backend/cmd"
	"log/slog"
)

func main() {
	// コマンドの実行
	// flow: コマンド -> コンフィグ -> ルーター
	err := cmd.Execute()
	if err != nil {
		slog.Warn(err.Error())
		return
	}
}
