## Install the protocol compiler plugins for Go using the following commands:

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## 创建 proto 文件

- 如文件夹下的 sorter.proto

## 生成 pb.go 文件

```shell
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative sorter/sorter.proto
```


## case 运行

先运行server端

```shell
go run server.go
```

在运行 client 端

```shell
go run client.go
```