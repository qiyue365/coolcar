package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/qiyue365/coolcar/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	log.SetFlags(log.Default().Flags() | log.Lshortfile)
}

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("127.0.0.1:11223", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("dial failed:", err)
	}
	defer conn.Close()

	client := pb.NewTripServiceClient(conn)
	res, err := client.GetTrip(ctx, &pb.GetTripRequest{Id: uuid.New().String()})
	if err != nil {
		log.Fatal("invoke GetTrip() failed:", err)
	}
	log.Println(res)
}
