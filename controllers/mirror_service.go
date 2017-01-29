package controllers

import (
	"errors"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"github.com/mirrorhub-io/platform/models"
	"golang.org/x/net/context"
)

type MirrorServiceServer struct {
}

func (m *MirrorServiceServer) Get(ctx context.Context,
	request *pb.MirrorGetRequest) (*pb.MirrorGetResponse, error) {
	if err := contactAuth(ctx); err != nil {
		return nil, err
	}
	mirrors := make([]*pb.Mirror, 0)
	mirrors = models.MirrorList(10, 0).ToProto()
	return &pb.MirrorGetResponse{
		Mirrors: mirrors,
	}, nil
}

func (m *MirrorServiceServer) Find(ctx context.Context,
	request *pb.MirrorFindRequest) (*pb.Mirror, error) {
	mirror, err := AuthMirror(ctx)
	if err != nil {
		return nil, err
	}
	return mirror.ToProto(), nil
}

func (m *MirrorServiceServer) Create(ctx context.Context, mirror *pb.Mirror) (*pb.Mirror, error) {
	if err := contactAuth(ctx); err != nil {
		return nil, err
	}
	x := models.MirrorFromProto(mirror)
	models.Connection().Create(&x)
	return x.ToProto(), nil
}

func contactAuth(ctx context.Context) error {
	contact, _, err := AuthContact(ctx)
	if err != nil {
		return err
	}
	if contact == nil {
		return errors.New("Unauthorized")
	}
	return nil
}
