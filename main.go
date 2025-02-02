// main.go
package main

import (
	"fmt"
	"os"

	"RClog/logger"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to create logger:", err)
		return
	}

	loggerInstance, err := logger.NewLogger(logger.INFO, os.Stdout, file)
	if err != nil {
		fmt.Println("Failed to create logger:", err)
		return
	}

	loggerInstance.Info("This is an info message")
	loggerInstance.Warn("This is a warning message")
	loggerInstance.Error("This is an error message")
}
