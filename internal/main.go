package main

import (
	"github.com/oogway93/golangArchitecture/internal/server"
	"github.com/oogway93/golangArchitecture/internal/server/http/handlers/shop"
)

func main() {
	server := new(server.Server)
	server.Run("8080", handlerShop.HandlerRoutes())
}

