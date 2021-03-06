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

func StartApi(addr, port string) {
	bind_uri := addr + ":" + port
	log.Info("Starting API with URI: " + bind_uri)
	lis, err := net.Listen("tcp", bind_uri)
	if err != nil {
		log.Fatal(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMirrorServiceServer(grpcServer, new(MirrorServiceServer))
	pb.RegisterContactServiceServer(grpcServer, new(ContactServiceServer))
	pb.RegisterServiceServiceServer(grpcServer, new(ServiceServiceServer))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func StartGateway(addr, port, api string) {
	bind_uri := addr + ":" + port
	log.Info("Starting Gateway with URI: " + bind_uri)
	log.Info("Expected API addr: " + api)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	mopts := []grpc.DialOption{grpc.WithInsecure()}
	pb.RegisterMirrorServiceHandlerFromEndpoint(ctx, mux, api, mopts)
	pb.RegisterContactServiceHandlerFromEndpoint(ctx, mux, api, mopts)
	pb.RegisterServiceServiceHandlerFromEndpoint(ctx, mux, api, mopts)
	http.ListenAndServe(bind_uri, mux)
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
