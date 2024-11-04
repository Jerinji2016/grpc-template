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

## Authentication

JWT is used authentication and is defined in `pkg/auth/jwt.go`. To start add jwt dependency with:

```sh
go get github.com/golang-jwt/jwt/v4
```
