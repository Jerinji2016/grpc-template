# gRPC Template 🚀

Template for gRPC backend project.

## Getting Started 🏁

1. Clone project with

    ```sh
    git clone git@github.com:Jerinji2016/grpc-template.git
    ```

2. To rename module, change the module in `go.mod` and replace all the imports respective in other files

3. Run `git mod tidy`

4. Create a `.env` file with envorinment variables with schema below

5. Start serving with `go run src/cmd/server.go`.

## Environment Schema 🍁

```dotenv
PORT=50051

JWT_SECRET=your_secret_key

DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=admin
DB_PASSWORD=secret
DB_NAME=grpc_sample
```

## Authentication 🔑

JWT is used authentication and is defined in `src/pkg/auth/jwt.go`. To start add jwt dependency with:

```sh
go get github.com/golang-jwt/jwt/v4
```

## Logger 🪵

Logger is implemented with [logrus](https://pkg.go.dev/github.com/sirupsen/logrus) package. To start add logrus dependency with:

```sh
go get github.com/sirupsen/logrus
```

Use logger with:

```go
//  info log
logger.infoLog("FYI %v!", info)

//  debug log
logger.DebugLog("Hello %v!", variable)

//  error log
logger.ErrorLog("Error: %v", err)

//  warning log
logger.WarnLog("Careful: %v", warning)
```

***

## Database 🗄️

Database is implemented with [postgresql](https://www.postgresql.org/) with [gorm](https://gorm.io/).

Start by installing the `gorm` & `postgres` as this [documentation](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL):

```sh
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

Define you schemas in `src/internal/models/` and then add the schema instance to `schema` defined on top in `db.go`.

```go
schemas := []interface{}{
    &models.Post{},
    //  other models
}
```

Make sure postgres services are running and credentials mentioned in .env is setup.

## Up Next ⏭️

- Connecting to DB
- Working with DB
- Working with streams
- Handling proto buffers for multi repo projects
