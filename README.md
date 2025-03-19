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

go-grpc-middlewareがmod tidyでinstallされない
> go get github.com/grpc-ecosystem/go-grpc-middleware
```

###  gRPC code 生成
```
protoファイル作成後
> protoc -I. --go_out=. --go-grpc_out=. proto/*.proto
validateする場合
protoc-gen-validateのパスは適宜合わせる
> protoc -I. -I=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.2.1 --go_out=. --validate_out="lang=go:." --go-grpc_out=. proto/*.proto
```
