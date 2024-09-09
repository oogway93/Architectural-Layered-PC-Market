package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	repositoryPostgres "github.com/oogway93/golangArchitecture/internal/repository/postgres"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/auth"
	repositoryPostgresShop "github.com/oogway93/golangArchitecture/internal/repository/postgres/shop"
	repositoryPostgresUser "github.com/oogway93/golangArchitecture/internal/repository/postgres/user"
	repositoryRedis "github.com/oogway93/golangArchitecture/internal/repository/redis"

	// "github.com/oogway93/golangArchitecture/internal/server/serverHTTP"
	"github.com/oogway93/golangArchitecture/internal/server/serverHTTP"
	HTTP "github.com/oogway93/golangArchitecture/internal/server/serverHTTP"
	// handlerShopCategory "github.com/oogway93/golangArchitecture/internal/server/serverHTTP/handler/shop/category"
	serviceAuth "github.com/oogway93/golangArchitecture/internal/service/auth"
	serviceShop "github.com/oogway93/golangArchitecture/internal/service/shop"
	serviceUser "github.com/oogway93/golangArchitecture/internal/service/user"
	// "github.com/oogway93/golangArchitecture/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file",
			err)
	}
	PORT := os.Getenv("PORT")
	DB_PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db := repositoryPostgres.DatabaseConnection(repositoryPostgres.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     DB_PORT,
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMode"),
	})

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Order{}, &models.Delivery{}, &models.OrderItem{})

	cache, err := repositoryRedis.New()
	if err != nil {
		log.Fatalf("Error initializaing redis")
	}
	defer cache.Close()

	// Category
	categoryRepo := repositoryPostgresShop.NewRepositoryCategoryShop(db)
	categoryService := serviceShop.NewServiceShopCategory(categoryRepo, cache)
	// categoryHandler := handlerShopCategory.NewCategoryShopHandler(categoryService)

	productRepo := repositoryPostgresShop.NewRepositoryProductShop(db)
	productService := serviceShop.NewServiceShopProduct(productRepo, cache)
	// categoryHandler := handlerShopCategory.NewCategoryShopHandler(categoryService)
	// categoryHandler := serverHTTP.SetupRouter(categoryService)

	orderRepo := repositoryPostgresShop.NewRepositoryOrderShop(db)
	orderService := serviceShop.NewServiceShopOrder(orderRepo, cache)
	// categoryHandler := handlerShopCategory.NewCategoryShopHandler(categoryService)
	// categoryHandler := serverHTTP.SetupRouter(categoryService)

	userRepo := repositoryPostgresUser.NewRepositoryUser(db)
	userService := serviceUser.NewServiceUser(userRepo, cache)
	// categoryHandler := handlerShopCategory.NewCategoryShopHandler(categoryService)
	// categoryHandler := serverHTTP.SetupRouter(categoryService)

	authRepo := repositoryPostgresAuth.NewRepositoryAuth(db)
	authService := serviceAuth.NewServiceAuth(authRepo, cache)
	// categoryHandler := handlerShopCategory.NewCategoryShopHandler(categoryService)
	// categoryHandler := serverHTTP.SetupRouter(categoryService)

	router := serverHTTP.SetupRouter(categoryService, productService, orderService, userService, authService)

	server := new(HTTP.Server)
	if err := server.Run(PORT, router); err != nil {
		log.Fatal("Some errors in initialization routes",
			err)
	}
}
