
$ go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
$ mkdir openapi
$ oapi-codegen -generate types,client -o openapi/client.go openapi.yml