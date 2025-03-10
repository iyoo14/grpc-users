# grpc-users

# 準備
## 環境設定
```
> brew install protobuf

https://grpc.io/docs/languages/go/quickstart/
> go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

bshrc zhrcファイルに
> export PATH="$PATH:$(go env GOPATH)/bin"
> protoc-gen-go --version
> protoc-gen-go-grpc --version

validateする場合
> go install github.com/envoyproxy/protoc-gen-validate@latest
```

###  gRPC code 生成
protoファイル作成後
> protoc -I. --go_out=. --go-grpc_out=. proto/*.proto

