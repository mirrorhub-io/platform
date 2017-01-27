package controllers

import (
	"errors"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"github.com/mirrorhub-io/platform/models"
	"golang.org/x/net/context"
)

type mirrorServiceServer struct {
}

func (m *mirrorServiceServer) Get(ctx context.Context, request *pb.MirrorGetRequest) (*pb.MirrorGetResponse, error) {
	_, err := auth(ctx)
	mirrors := make([]*pb.Mirror, 0)
	if err == nil {
		mirrors = models.MirrorList(10, 0).ToProto()
	}
	return &pb.MirrorGetResponse{
		Mirrors: mirrors,
	}, err
}

func (m *mirrorServiceServer) Find(ctx context.Context, request *pb.MirrorFindRequest) (*pb.Mirror, error) {
	return &pb.Mirror{}, nil
}

func (m *mirrorServiceServer) Create(ctx context.Context, mirror *pb.Mirror) (*pb.Mirror, error) {
	x := models.MirrorFromProto(mirror)
	models.Connection().Create(&x)
	return x.ToProto(), nil
}

func auth(ctx context.Context) (*models.Mirror, error) {
	mirror := AuthMirror(ctx)
	if mirror != nil {
		return mirror, nil
	}
	return mirror, errors.New("Unauthorized")
}
