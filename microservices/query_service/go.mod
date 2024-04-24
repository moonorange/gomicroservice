module query_service

go 1.22.2

require (
	connectrpc.com/connect v1.16.1
	github.com/protogo v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/net v0.23.0
	golang.org/x/sync v0.6.0
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/protogo => ./proto_go
