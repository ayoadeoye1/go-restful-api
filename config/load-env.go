package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var EnvFile map[string]string

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Error determining directory: %v", err)
	}

	envPath := filepath.Join(dir, "../.env")

	EnvFile, err = godotenv.Read(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
