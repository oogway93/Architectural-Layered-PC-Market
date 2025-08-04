package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/oogway93/golangArchitecture/internal/adapter/config"
	"github.com/oogway93/golangArchitecture/internal/adapter/logger"
	repositoryPostgres "github.com/oogway93/golangArchitecture/internal/core/repository/postgres"
	repositoryPostgresAuth "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/auth"
	"github.com/oogway93/golangArchitecture/internal/core/repository/postgres/models"
	repositoryPostgresShop "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/shop"
	repositoryPostgresUser "github.com/oogway93/golangArchitecture/internal/core/repository/postgres/user"
	repositoryRedis "github.com/oogway93/golangArchitecture/internal/core/repository/redis"

	Server "github.com/oogway93/golangArchitecture/internal/core/server"

	serviceAuth "github.com/oogway93/golangArchitecture/internal/core/service/auth"
	serviceShop "github.com/oogway93/golangArchitecture/internal/core/service/shop"
	serviceUser "github.com/oogway93/golangArchitecture/internal/core/service/user"
)

func main() {
	env := flag.String("env", "development", "env's status")
	flag.Parse()
	gin.SetMode(gin.DebugMode)
	APP_ENV := *env
	config, err := config.New(APP_ENV)
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", APP_ENV)

	db, err := repositoryPostgres.DatabaseConnection(repositoryPostgres.Config{
		Username: config.DB.User,
		Password: config.DB.Password,
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		DBName:   config.DB.Name,
		SSLMode:  config.DB.SSLMode,
	})
	if err != nil {
		slog.Warn("DB connection isn`t successful: %s", err)
	}
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Order{}, &models.Delivery{}, &models.OrderItem{})
	slog.Info("Successfully migrated the database")

	addr := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	cache, err := repositoryRedis.New(repositoryRedis.Config{
		Addr:       addr,
		Password:   config.Redis.Password,
		Expiration: time.Duration(config.Redis.Expiration) * time.Minute,
	})
	if err != nil {
		slog.Error("Error initializing cache connection", "error", err)
	}
	defer cache.Close()

	slog.Info("Successfully connected to the cache server")

	// Category
	categoryRepo := repositoryPostgresShop.NewRepositoryCategoryShop(db)
	categoryService := serviceShop.NewServiceShopCategory(categoryRepo, cache)

	// Product
	productRepo := repositoryPostgresShop.NewRepositoryProductShop(db)
	productService := serviceShop.NewServiceShopProduct(productRepo, cache)

	// Order
	orderRepo := repositoryPostgresShop.NewRepositoryOrderShop(db)
	orderService := serviceShop.NewServiceShopOrder(orderRepo, cache)

	// User
	userRepo := repositoryPostgresUser.NewRepositoryUser(db)
	userService := serviceUser.NewServiceUser(userRepo, cache)

	// Auth
	authRepo := repositoryPostgresAuth.NewRepositoryAuth(db)
	authService := serviceAuth.NewServiceAuth(authRepo, cache)

	router := Server.SetupRouter(config.HTTP, categoryService, productService, orderService, userService, authService)
	server := new(Server.Server)
	if err := server.Run(config, router); err != nil {
		slog.Error("Some errors in starting Server", "error", err)
	}
}
