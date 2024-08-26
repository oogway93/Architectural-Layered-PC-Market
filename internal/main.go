package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	repositoryPostgres "github.com/oogway93/golangArchitecture/internal/repository/postgres"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	HTTP "github.com/oogway93/golangArchitecture/internal/server/http"

	"github.com/oogway93/golangArchitecture/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file",
			err)
	}
	PORT := os.Getenv("PORT")
	DB_PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db:= repositoryPostgres.DatabaseConnection(repositoryPostgres.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     DB_PORT,
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMode"),
	})
	db.Table("users").AutoMigrate(&models.User{})
	db.Table("categories").AutoMigrate(&models.Category{})

	repo := repositoryPostgres.NewRepository(db)
	service := service.NewService(repo)
	router := HTTP.SetupRouter(service)

	server := new(HTTP.Server)
	if err := server.Run(PORT, router); err != nil {
		log.Fatal("Some errors in initialization routes",
			err)
	}
}
