package proto

//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/indiealistic/bethacon/proto/status/status.proto
//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/indiealistic/bethacon/proto/health/health.proto
//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/indiealistic/bethacon/proto/common/error_response.proto
//go:generate protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/indiealistic/bethacon/proto/common/types.proto

//go:generate protoc -I $GOPATH/src/github.com/prysmaticlabs/ethereumapis --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $GOPATH/src/github.com/indiealistic/bethacon/proto/bethacon/bethacon.proto
