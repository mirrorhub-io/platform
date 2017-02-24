package controllers

import (
	"errors"
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
	pb.RegisterServiceServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", mopts)
	go http.ListenAndServe(":8080", mux)
	lis, _ := net.Listen("tcp", ":9000")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMirrorServiceServer(grpcServer, new(MirrorServiceServer))
	pb.RegisterContactServiceServer(grpcServer, new(ContactServiceServer))
	pb.RegisterServiceServiceServer(grpcServer, new(ServiceServiceServer))
	grpcServer.Serve(lis)
}

func AuthContact(ctx context.Context) (*models.Contact, string, error) {
	md, _ := metadata.FromContext(ctx)
	if md["contacttoken"] == nil {
		log.Debug("Token missing")
		return nil, "", errors.New("Token missing")
	}
	contact, err := models.AuthContactWithToken(
		md["contacttoken"][0],
	)
	if err != nil {
		return nil, "", err
	}
	if contact == nil {
		return nil, "", errors.New("Contact not resolvable.")
	}
	return contact, md["contacttoken"][0], nil
}

func AuthMirror(ctx context.Context) (*models.Mirror, error) {
	md, _ := metadata.FromContext(ctx)
	if md["clienttoken"] == nil {
		return nil, errors.New("Client token missing")
	}
	mirror := &models.Mirror{}
	models.Connection().Where(
		"client_token = ?",
		md["clienttoken"][0],
	).First(&mirror)
	if models.Connection().NewRecord(mirror) {
		return nil, errors.New("No mirror found.")
	}
	return mirror, nil
}
