package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	bethaconproto "github.com/indiealistic/bethacon/proto/bethacon"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bethaConService := bethaconproto.NewBethaConServiceClient(conn)

	resp, err := bethaConService.Ping(context.Background(), &empty.Empty{})
	fmt.Println("resp", resp)
	fmt.Println("err", err)
}
