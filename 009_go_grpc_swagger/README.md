# GRPC with Swagger

## Modern Way Build for buf

```shell
$ brew install bufbuild/buf/buf
```

```shell
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### Old Way 

```shell
$ protoc \  
-I proto \  
-I third_party/grpc-gateway/ \  
-I third_party/googleapis \  
--go_out=plugins=grpc,paths=source_relative:./api/v1 \  
--grpc-gateway_out=paths=source_relative:./api/v1 \  
--openapiv2_out=third_party/OpenAPI/ \  
api/v1/*.proto
```

### New Way 

```shell
$ protoc \  
-I proto \  
-I third_party/grpc-gateway/ \  
-I third_party/googleapis \  
-I vendor \
--go_out=plugins=grpc,paths=source_relative:./api/v1 \  
--grpc-gateway_out=paths=source_relative:./api/v1 \  
--openapiv2_out=third_party/OpenAPI/ \  
api/v1/*.proto
```

## 이슈 관련 링크 자료 정리 

- [buf generate in usage documentation](https://github.com/grpc-ecosystem/grpc-gateway/issues/2039)

## 참조 링크 

> [GRPC with Swagger](https://medium.com/@pointgoal/grpc-how-to-add-swagger-ui-on-grpc-466e5fd71097)   
> [Buf Build with Swagger](https://medium.com/@vchitai/using-buf-build-to-generate-your-grpc-codes-44e1811d5291)   
> [Buf Build](https://docs.buf.build/installation)   