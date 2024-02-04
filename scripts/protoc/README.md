## generate
```
protoc --go_out=.  --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto
```

```
protoc -I. --go_out=internal/infrastructure/grpc --go-grpc_out=internal/infrastructure/grpc api/*.proto
```

## delete
```
rm greeting/greeting*pb*
```