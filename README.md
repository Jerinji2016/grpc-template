# gRPC Template 🚀

Template for gRPC backend project.

## Getting Started 🏁

1. Clone project with

    ```sh
    git clone git@github.com:Jerinji2016/grpc-template.git
    ```

2. To rename module, change the module in `go.mod` and replace all the imports respective in other files

3. Run `git mod tidy`

4. Start serving with `go run cmd/server.go`.

## Authentication 🔑

JWT is used authentication and is defined in `pkg/auth/jwt.go`. To start add jwt dependency with:

```sh
go get github.com/golang-jwt/jwt/v4
```

## Logger 🪵

Logger is implemented with logrus package. To start add logrus dependency with:

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

## Up Next ⏭️

- Environment files
- Connecting to DB
- Working with DB
- Working with streams
- Handling proto buffers for multi repo projects
