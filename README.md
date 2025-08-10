# Golang Architectural Layered PC Market

## Project's stack

* Gin
* Postgres
* JWT
* Hashing password
* Environments: production(docker) and development(locally)
* Docker and Docker Compose
* Cookies
* Authorization
* Nginx(Proxy)
* Redis
* Logging
* Middlewares
* API and "Human" parts
* HTTP/2 and TLS

### How to start the project(Locally|Development):

1. Create the file .env.development in root of project:

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

TLS_CERT_PATH=internal/adapter/certifications/cert.pem
TLS_KEY_PATH=internal/adapter/certifications/key.pem
TEMPLATES_PATH=internal/core/server/serverHTTP/static/templates/shop

LOG_FILE_PATH=logs/application.log
```
2. Create cert and key for TLS connection(HTTPS)
```zsh
    cd internal/adapter/certifications
        
    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
``` 
3. Start the app
```zsh 
    make rundev
```

4. Let's check out our connection to the golang's server in a browser by an URL:
```
    https://localhost:8000/
```

### How to start the project(Production)
1. Create the file .env.production in root of project:
```golang
DB_PORT=5432
DB_HOST=postgres
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSLMode=disable
SECRET=8DxZ4jzMmA

REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=redis
REDIS_EXPIRATION=1

HTTP_URL=app
HTTP_PORT=8000
APP_NAME=golangArchitecture

TLS_CERT_PATH=/app/internal/adapter/certifications/cert.pem
TLS_KEY_PATH=/app/internal/adapter/certifications/key.pem
TEMPLATES_PATH=/app/internal/core/server/serverHTTP/static/templates/shop

LOG_FILE_PATH=/app/logs/application.log
```
2. Create cert and key for TLS connection(HTTPS)
```zsh
   cd internal/adapter/certifications
        
    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```
3. Run the project
```zsh
    make runprod
```

#### RUNNING TESTS
1. Create the file ".env.test" in root with this data:
```golang
TEST_DB_PORT=5432
TEST_DB_HOST=localhost
TEST_DB_USERNAME=postgres
TEST_DB_PASSWORD=postgres
TEST_DB_NAME=testdb
TEST_DB_SSLMode=disable
```
2. Start tests by the command below
```zsh
make tests
```
