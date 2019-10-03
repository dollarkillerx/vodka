生成grpc 代码
``` 
protoc --go_out=plugins=grpc:hello h1.proto
```

生成序列化 代码
``` 
protoc --go_out=. h1.proto
```