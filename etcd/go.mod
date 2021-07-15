module github.com/golang-learning/etcd

go 1.15

require (
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.3.0
	go.etcd.io/etcd v3.3.25+incompatible
	go.etcd.io/etcd/client/v3 v3.5.0 // indirect
	go.uber.org/zap v1.18.1 // indirect
	google.golang.org/grpc v1.39.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
