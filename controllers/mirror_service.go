package controllers

import (
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"github.com/mirrorhub-io/platform/models"
	"golang.org/x/net/context"
)

type mirrorServiceServer struct {
}

func (m *mirrorServiceServer) Get(ctx context.Context, request *pb.MirrorGetRequest) (*pb.MirrorGetResponse, error) {
	return &pb.MirrorGetResponse{
		Mirrors: models.MirrorList(10, 0).ToProto(),
	}, nil
}

func (m *mirrorServiceServer) Create(ctx context.Context, mirror *pb.Mirror) (*pb.Mirror, error) {
	x := models.MirrorFromProto(mirror)
	models.Connection().Create(&x)
	return x.ToProto(), nil
}
