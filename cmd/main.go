package main

import (
	"github.com/mirjalilova/auth-service-blacklist/internal/config"
	"github.com/mirjalilova/auth-service-blacklist/pkg/app"
)

func main() {
	// Load configuration from environment variables or a file.
	cfg := config.Load()

	// Run the application with the loaded configuration.
	app.Run(&cfg)
}
