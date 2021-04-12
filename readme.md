生成protobugf
<br>
protoc --go-grpc_out=. *.proto
<br>
protoc --go_out=. *.proto

需要将 
,omitempty
替换为 
空



protoc 
说明文档
https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go
下载地址
https://github.com/protocolbuffers/protobuf/releases/tag/v3.15.8

protoc go插件
说明文档
https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go
安装指南
https://developers.google.com/protocol-buffers/docs/gotutorial
