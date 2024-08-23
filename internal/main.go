package main

import (
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	repositoryPostgres "github.com/oogway93/golangArchitecture/internal/repository/postgres"
	"github.com/oogway93/golangArchitecture/internal/server/http"
	// "github.com/oogway93/golangArchitecture/internal/server/http/handler/shop"

	handlerUser "github.com/oogway93/golangArchitecture/internal/server/http/handler/user"
	"github.com/oogway93/golangArchitecture/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file",
			err)
	}
	PORT := os.Getenv("PORT")

	db, err := repositoryPostgres.NewPostgresDB(repositoryPostgres.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMode"),
	})

	if err != nil {
		log.Fatal("Failed to initialized db",
			err)
	}

	repo := repositoryPostgres.NewRepository(db)
	service := service.NewService(repo)
	// handlersProduct :=  handlerShop.NewCategoryShopHandler(service)
	handlersUser :=  handlerUser.NewUserHandler(service)
	
	server := new(http.Server)
	if err := server.Run(PORT, handlersUser.UserHandlerRoutes()); err != nil {
		log.Fatal("Some errors in initialization routes",
			err)
	}
}

