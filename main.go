package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"dice/pb"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("start gRPC Client.")

	client := pb.NewDiceServiceClient(conn)

	res, err := client.RollDice(
		context.Background(),
		&pb.RollDiceRequest{
			MinNum: 1,
			MaxNum: 6,
		},
	)
	if err != nil {
		panic(err)
	}

	var result int32 = res.GetResult()
	fmt.Println("サイコロの値:", result)
}
