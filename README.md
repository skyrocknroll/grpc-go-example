### gRPC Golang example
```bash
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```
### For new version 
```
protoc -I protos/otlp --go_out=protos/otlp --go_opt=paths=source_relative   --go-grpc_out=protos/otlp --go-grpc_opt=paths=source_relative   protos/otlp/*.proto
```
