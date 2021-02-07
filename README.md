# Authenticator Fetcher

Authenticator and Fetcher application.

## Installing / Getting started

A quick introduction of the minimal setup to running the app.

- Authenticator 
    ```shell
    cd authenticator
    cp .env.example .env
    go mod tidy
    go run main.go
    ```

- Fetcher
    ```shell
    cd fetcher
    cp .env.example .env
    yarn install
    node main.js
    ```

- Docker 
    ```shell
    docker-compose up
    ```

## Developing
### Built With
- Authenticator
    1. Golang
    2. Postgresql
    3. Gin Gonic

- Fetcher
    1. Node.js
    2. Express.js
    3. Redis

(Optional) Docker

### Setting up Dev

Here's a brief intro about what a developer must do in order to start developing
the project further:

```shell
git clone https://github.com/zainokta/authenticator-fetcher
cd authenticator-fetcher/
cd authenticator/
go mod tidy

cd ../fetcher
yarn install
```

(Optional) Docker
```shell
docker-compose up --build
```

Fill up the .env file with your local config (must)

## Configuration

- Authenticator
    ```text
    APP_PORT= //Application port

    GIN_MODE= //Gin mode "release" or "debug"

    DB_DRIVER= //Database driver, in this project I use postgre, so it must be "postgres"

    DB_DATABASE= //Your own database name

    DB_HOST= //Database host, local or remote or even docker image

    DB_PORT= //Database Port, for postgres it's usuall 5432

    DB_USERNAME= //Database username

    DB_PASSWORD= //Database password

    JWT_SECRET= //JWT salt for security, make sure it's same with fether app
    ```
    
- Fetcher 
    ```text
    APP_PORT= //Application port

    REDIS_HOST= //Redis host i.e (redis://cache)

    BASE_STEIN_URL = https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list
    
    BASE_CURRENCY_CONVERTER_URL = https://free.currconv.com/api/v7/convert?q=USD_IDR&compact=ultra&apiKey=547807587094635f1a50

    JWT_SECRET= //JWT salt, must be same with Authenticator app

    TIMEOUT= //Axios timeout
    ```

## Database

- For Authenticator app, I use postgresql database, if you using docker, it will automatically pull the latest version of postgres
- For Fetcher app, it's only using redis for cache stein API and currency converter
