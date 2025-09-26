package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	KeycloakURL  string
	ClientID     string
	ClientSecret string
	Realm        string
	RedirectURI  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		KeycloakURL:  os.Getenv("KEYCLOAK_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Realm:        os.Getenv("REALM"),
		RedirectURI:  os.Getenv("REDIRECT_URI"),
	}
}
