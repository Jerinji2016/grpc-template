## Godoc for Go Code

The standard way to document Go code is by using **Godoc**, which is built into the Go ecosystem. By adding comments to your Go code in a specific format, Godoc can generate HTML documentation.

#### Use comments for Documentation as below

```go
// UserService handles user-related operations.
type UserService struct {
    // some fields...
}

// CreateUser creates a new user in the system.
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    // implementation...
}
```

Generate documentation by running

```sh
godoc -http=:6060
```

This serves the documentation locally at http://localhost:6060. You can browse the documentation for all Go packages, including your own.

***

### Swagger (OpenAPI) for REST APIs via gRPC-Gateway

1. **Install gRPC-Gateway**

```sh
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```

2. **Update the proto file** to include HTTP annotations:
   
```protobuf
syntax = "proto3";

package userapi;

import "google/api/annotations.proto";  // Import the annotations proto

service UserService {
    rpc CreateUser(CreateUserRequest) returns (CreateUserResponse) {
        option (google.api.http) = {
            post: "/v1/users"
            body: "user"
        };
    }
    rpc GetUser(GetUserRequest) returns (GetUserResponse) {
        option (google.api.http) = {
            get: "/v1/users/{id}"
        };
    }
}
```

3. **Generate the gRPC-Gateway and OpenAPI spec**:
   
```sh
protoc -I. \
  -I$GOPATH/src \
  --grpc-gateway_out ./gen --grpc-gateway_opt logtostderr=true \
  --openapiv2_out ./swagger --openapiv2_opt logtostderr=true \
  api/user.proto
```

4. **Serve the Swagger UI**:
   
   You can use **Swagger UI** to serve the OpenAPI spec interactively. Swagger UI can be set up to serve the OpenAPI file (swagger.json or swagger.yaml) and display the API documentation in an interactive web interface.

***