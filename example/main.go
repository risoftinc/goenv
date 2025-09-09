package main

import (
	"fmt"
	"log"
	"time"

	"github.com/risoftinc/goenv"
)

func main() {
	// Example 1: Load from .env file (key-value format)
	fmt.Println("=== Loading from .env file ===")
	err := goenv.LoadEnv("config.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Get values using dot notation for nested keys
	appName := goenv.GetEnvString("APP_NAME", "DefaultApp")
	appVersion := goenv.GetEnvString("APP_VERSION", "1.0.0")
	debug := goenv.GetEnvBool("DEBUG", false)
	port := goenv.GetEnvInt("PORT", 8080)

	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("App Version: %s\n", appVersion)
	fmt.Printf("Debug: %t\n", debug)
	fmt.Printf("Port: %d\n", port)

	// Example 2: Load from JSON file
	fmt.Println("\n=== Loading from JSON file ===")
	err = goenv.LoadEnv("config.json")
	if err != nil {
		log.Printf("Error loading JSON file: %v", err)
	}

	// Get nested values using dot notation
	dbHost := goenv.GetEnvString("database.host", "localhost")
	dbPort := goenv.GetEnvInt("database.port", 5432)
	dbName := goenv.GetEnvString("database.name", "mydb")
	timeout := goenv.GetEnvFloat64("timeout", 30.0)

	fmt.Printf("Database Host: %s\n", dbHost)
	fmt.Printf("Database Port: %d\n", dbPort)
	fmt.Printf("Database Name: %s\n", dbName)
	fmt.Printf("Timeout: %.1f\n", timeout)

	// Example 3: Load from YAML file
	fmt.Println("\n=== Loading from YAML file ===")
	err = goenv.LoadEnv("config.yaml")
	if err != nil {
		log.Printf("Error loading YAML file: %v", err)
	}

	// Get nested values using dot notation
	apiHost := goenv.GetEnvString("api.host", "localhost")
	apiPort := goenv.GetEnvInt("api.port", 3000)
	apiTimeout := goenv.GetEnvInt("api.timeout", 30)

	fmt.Printf("API Host: %s\n", apiHost)
	fmt.Printf("API Port: %d\n", apiPort)
	fmt.Printf("API Timeout: %d\n", apiTimeout)

	// Example 4: Using GetEnvNested for dot notation
	fmt.Println("\n=== Using GetEnvNested ===")
	// This will look for DB_HOST environment variable
	dbHostNested := goenv.GetEnvNested("db.host", "localhost")
	dbPortNested := goenv.GetEnvNested("db.port", 5432)

	fmt.Printf("DB Host (nested): %s\n", dbHostNested)
	fmt.Printf("DB Port (nested): %d\n", dbPortNested)

	// Example 5: Load with specific format
	fmt.Println("\n=== Loading with specific format ===")
	err = goenv.LoadEnvWithFormat(goenv.FormatJSON, "config.json")
	if err != nil {
		log.Printf("Error loading JSON with specific format: %v", err)
	}

	// Example 6: Generic GetEnv function
	fmt.Println("\n=== Using generic GetEnv ===")
	genericString := goenv.GetEnv("app.name", "DefaultApp")
	genericInt := goenv.GetEnv("app.port", 8080)
	genericBool := goenv.GetEnv("app.debug", false)
	genericDuration := goenv.GetEnv("timeout", 30*time.Second)

	fmt.Printf("Generic String: %s\n", genericString)
	fmt.Printf("Generic Int: %d\n", genericInt)
	fmt.Printf("Generic Bool: %t\n", genericBool)
	fmt.Printf("Generic Duration: %v\n", genericDuration)

	// Example 7: Duration support
	fmt.Println("\n=== Duration Support ===")
	timeoutData := goenv.GetEnvDuration("TIMEOUT", 30*time.Second)
	retryInterval := goenv.GetEnvDuration("RETRY_INTERVAL", 5*time.Minute)
	cleanupInterval := goenv.GetEnvDuration("CLEANUP_INTERVAL", 1*time.Hour)

	fmt.Printf("Timeout: %v\n", timeoutData)
	fmt.Printf("Retry Interval: %v\n", retryInterval)
	fmt.Printf("Cleanup Interval: %v\n", cleanupInterval)
}
