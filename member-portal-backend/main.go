package main

import (
	"github.com/kstm-su/Member-Portal/backend/cmd"
)

func main() {
	// Initialize a Router
	err := cmd.Execute()
	if err != nil {
		err.Error()
		return
	}
}
