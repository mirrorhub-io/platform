package controllers

import (
	models "../models"
	pb "./proto"
	log "github.com/Sirupsen/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"time"
)

type mirrorServiceServer struct {
}

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

func (m *mirrorServiceServer) Get(ctx context.Context, request *pb.MirrorGetRequest) (*pb.MirrorGetResponse, error) {
	log.Info("test")
	return &pb.MirrorGetResponse{
		Mirrors: []*pb.Mirror{
			&pb.Mirror{
				Name:        "Moo",
				OnlineSince: time.Now().Unix(),
			},
		},
	}, nil
}

func (m *mirrorServiceServer) Create(ctx context.Context, mirror *pb.Mirror) (*pb.Mirror, error) {
	x := models.Mirror{
		Name: mirror.Name,
		IPv4: mirror.Ipv4,
		IPv6: mirror.Ipv6,
	}
	models.Connection().Create(&x)
	return x.ToProto(), nil
}
