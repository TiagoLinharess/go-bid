## Go Bid

### What is?
Go Bid is an API where users can register their products and receive bids for them; the highest bidder wins the product, similar to eBay.

### Instalation

#### 1 - create a .env file like the example.env

```env
GOBID_DATABASE_PORT=5432
GOBID_DATABASE_NAME="gobid"
GOBID_DATABASE_USER="docker"
GOBID_DATABASE_PASSWORD="docker"
GOBID_DATABASE_HOST="localhost"
GOBID_CSRF_KEY="NhkEXjyS5ms3k7vNQ5fbk2Ffv0OIuQs6"
```

#### 2 - run the docker compose image to create the database

```shell
docker compose up -d
```

#### 3 - install the dependencies

```shell
go mod tidy
```

#### 4 - run the /terndotenv main.go to run the migrations using [tern](https://github.com/jackc/tern)

```shell
go run ./cmd/terndotenv
```

#### 5 - run the command below for dev mode using [air](https://github.com/air-verse/air)

```shell
air --build.cmd "go build -o ./bin/api.exe ./cmd/api" --build.bin "./bin/api.exe"
```