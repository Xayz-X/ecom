package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// server and database configuaraion
type Config struct {
	// server configs
	PublicHost string
	Port       string

	// database configs
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

// instanciate a new configuration, by read the env variables
func initConfig() Config {
	// load the environment variables
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "ecom"),
	}
}

// get the env variable with fallback value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
