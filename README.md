# Golang Architecture Project


## Project's stack

* Gin
* Postgres
* JWT
* Hashing password
* Environments: production(docker) and development(locally)
* Docker and Docker Compose
* Cookies
* Authorization
* Nginx
* Redis
* Logging
* Middlewares
* API and "Human" parts
* HTTP/2 and TLS

### How to start the project(Locally):

1. Make the file .env.development in root of project:

```golang
DB_PORT=5432
DB_HOST=localhost
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSLMode=disable
SECRET=auth-jwt-token

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=redis
REDIS_EXPIRATION=1

HTTP_URL=localhost
HTTP_PORT=8000
APP_NAME=golangArchitecture
TLS_CERT_PATH=.../golangArchitecture/internal/adapter/tls/server.crt
TLS_KEY_PATH=.../golangArchitecture/internal/adapter/tls/server.key

LOG_FILE_PATH=.../golangArchitecture/logs/application.log
```
2. Create server.crt and server.key for TLS connection
```zsh
    cd internal/adapter/tls
        
    openssl genrsa -out server.key 2048

    openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
``` 
3. Start the container postgres

``` golang
    go run cmd/main.go -env=development
```

4. Let's check a browser by an URL:
```
    http://localhost:8000/
```

### How to start the project(Production)
1. Make the file .env.production in root of project:
```golang
DB_PORT=5432
DB_HOST=golangArchitecture_postgres
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSLMode=disable
SECRET=8DxZ4jzMmA

REDIS_HOST=golangArchitecture_redis
REDIS_PORT=6379
REDIS_PASSWORD=redis
REDIS_EXPIRATION=1

HTTP_URL=app
HTTP_PORT=8000
APP_NAME=golangArchitecture
TLS_CERT_PATH=/app/internal/adapter/tls/server.crt
TLS_KEY_PATH=/app/internal/adapter/tls/server.key
TEMPLATES_PATH=/app/internal/core/server/serverHTTP/static/templates/shop

LOG_FILE_PATH=/app/logs/application.log
```
2. Create server.crt and server.key for TLS connection
```zsh
    cd internal/adapter/tls
        
    openssl genrsa -out server.key 2048

    openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```
3. Run the project
```zsh
    sudo docker compose --env-file .env.production up -d --build
```

