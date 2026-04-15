package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// colorizeLogLevel colorizes INFO and ERROR log levels in the message
func colorizeLogLevel(message string) string {
	// Define color formatters
	infoColor := color.New(color.FgBlack, color.BgGreen).SprintFunc()
	errorColor := color.New(color.FgWhite, color.BgRed).SprintFunc()

	// Replace INFO with colored version
	message = strings.ReplaceAll(message, "INFO", infoColor("INFO"))

	// Replace ERROR with colored version
	message = strings.ReplaceAll(message, "ERROR", errorColor("ERROR"))

	return message
}

func main() {
	fmt.Println("Saw Color Demo - INFO and ERROR colorization")
	fmt.Println("==============================================")
	fmt.Println()

	// Example log messages
	messages := []string{
		"2024-04-15 14:30:00 INFO Application started successfully",
		"2024-04-15 14:30:01 INFO User logged in: john.doe",
		"2024-04-15 14:30:05 ERROR Failed to connect to database",
		"2024-04-15 14:30:06 INFO Retrying connection...",
		"2024-04-15 14:30:07 ERROR Connection timeout after 5 seconds",
		"2024-04-15 14:30:10 INFO Using fallback configuration",
		"[2024-04-15T14:30:15Z] (stream-1) INFO Starting worker process",
		"[2024-04-15T14:30:16Z] (stream-2) ERROR Worker crashed unexpectedly",
		`{"timestamp": "2024-04-15T14:30:20Z", "level": "INFO", "message": "API request processed"}`,
		`{"timestamp": "2024-04-15T14:30:21Z", "level": "ERROR", "message": "Invalid request payload"}`,
	}

	fmt.Println("Sample log messages with colorization:")
	fmt.Println()

	for _, msg := range messages {
		fmt.Println(colorizeLogLevel(msg))
	}

	fmt.Println()
	fmt.Println("Note: INFO appears with green background and black text")
	fmt.Println("      ERROR appears with red background and white text")
}
