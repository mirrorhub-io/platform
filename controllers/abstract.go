package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"github.com/mirrorhub-io/platform/models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	pb.RegisterContactServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", mopts)
	go http.ListenAndServe(":8080", mux)
	lis, _ := net.Listen("tcp", ":9000")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMirrorServiceServer(grpcServer, new(MirrorServiceServer))
	pb.RegisterContactServiceServer(grpcServer, new(ContactServiceServer))
	grpcServer.Serve(lis)
}

func AuthContact(ctx context.Context) (*models.Contact, string) {
	md, _ := metadata.FromContext(ctx)
	log.Info(md)
	if md["contactemail"] == nil {
		return nil, ""
	}
	if md["contacttoken"] == nil {
		return nil, ""
	}
	return models.AuthContactWithToken(
		md["contactemail"][0],
		md["contacttoken"][0],
	)
}

func AuthMirror(ctx context.Context) *models.Mirror {
	md, _ := metadata.FromContext(ctx)
	log.Info(md)
	if md["clienttoken"] == nil {
		return nil
	}
	mirror := &models.Mirror{}
	models.Connection().Where(
		"client_token = ?",
		md["clienttoken"][0],
	).First(&mirror)
	if models.Connection().NewRecord(mirror) {
		return nil
	}
	return mirror
}
