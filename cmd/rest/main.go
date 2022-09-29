package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/qiyue365/coolcar/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	endpoint := "127.0.0.1:11223"
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := pb.RegisterTripServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatal("register grpc-gateway failed: ", err)
	}
	log.Fatal(http.ListenAndServe(":11224", mux))
}
