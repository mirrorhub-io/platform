package controllers

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func StartServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	mopts := []grpc.DialOption{grpc.WithInsecure()}
	pb.RegisterMirrorServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", mopts)
	go http.ListenAndServe(":8080", mux)
	lis, _ := net.Listen("tcp", ":9000")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMirrorServiceServer(grpcServer, new(mirrorServiceServer))
	grpcServer.Serve(lis)
}
