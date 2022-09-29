package main

import (
	"log"
	"net"

	"github.com/qiyue365/coolcar/pb"
	"github.com/qiyue365/coolcar/service/trip"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.Default().Flags() | log.Lshortfile)
}

func main() {
	lis, err := net.Listen("tcp", ":11223")
	if err != nil {
		log.Fatal("listen failed:", err)
	}
	srv := grpc.NewServer()
	pb.RegisterTripServiceServer(srv, &trip.Service{})
	log.Fatal("grpc serve failed:", srv.Serve(lis))
}
