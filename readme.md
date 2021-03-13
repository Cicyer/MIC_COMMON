
生成protobugf
<br>
protoc --go-grpc_out=. *.proto
<br>
protoc --go_out=. *.proto

需要将 
omit去掉empty"
替换为 
"v