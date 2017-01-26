package controllers

import (
	log "github.com/Sirupsen/logrus"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"github.com/mirrorhub-io/platform/models"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type mirrorServiceServer struct {
}

func (m *mirrorServiceServer) Get(ctx context.Context, request *pb.MirrorGetRequest) (*pb.MirrorGetResponse, error) {
	md, _ := metadata.FromContext(ctx)
	log.Info("Metadata: ")
	log.Info(md)
	return &pb.MirrorGetResponse{
		Mirrors: models.MirrorList(10, 0).ToProto(),
	}, nil
}

func (m *mirrorServiceServer) Create(ctx context.Context, mirror *pb.Mirror) (*pb.Mirror, error) {
	x := models.MirrorFromProto(mirror)
	models.Connection().Create(&x)
	return x.ToProto(), nil
}
