package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Container contains environment variables for the application, database, cache, token, and http server
type (
	Container struct {
		App   *App
		Redis *Redis
		DB    *DB
		HTTP  *HTTP
	}
	// App contains all the environment variables for the application
	App struct {
		Name    string
		Env     string
		LogPath string
	}

	// Redis contains all the environment variables for the cache service
	Redis struct {
		Host       string
		Port       string
		Password   string
		Expiration int
	}
	// Database contains all the environment variables for the database
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
	}
	// HTTP contains all the environment variables for the http server
	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
		TLSCertPath    string
		TLSKeyPath     string
		TemplatesPath  string
	}
)

// New creates a new container instance
func New(APP_ENV string) (*Container, error) {
	if APP_ENV == "production" {
		err := godotenv.Load(".env.production")
		if err != nil {
			return nil, err
		}
	} else {
		err := godotenv.Load(".env.development")
		if err != nil {
			return nil, err
		}
	}
	redis_expiration, _ := strconv.Atoi(os.Getenv("REDIS_EXPIRATION"))
	app := &App{
		Name:    os.Getenv("APP_NAME"),
		Env:     APP_ENV,
		LogPath: os.Getenv("LOG_FILE_PATH"),
	}

	redis := &Redis{
		Host:       os.Getenv("REDIS_HOST"),
		Port:       os.Getenv("REDIS_PORT"),
		Password:   os.Getenv("REDIS_PASSWORD"),
		Expiration: redis_expiration,
	}

	db := &DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	http := &HTTP{
		Env:         os.Getenv("APP_ENV"),
		URL:         os.Getenv("HTTP_URL"),
		Port:        os.Getenv("HTTP_PORT"),
		TLSCertPath: os.Getenv("TLS_CERT_PATH"),
		TLSKeyPath:  os.Getenv("TLS_KEY_PATH"),
		TemplatesPath: os.Getenv("TEMPLATES_PATH"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	return &Container{
		app,
		redis,
		db,
		http,
	}, nil
}
