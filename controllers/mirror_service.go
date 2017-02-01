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

func (m *MirrorServiceServer) Connect(ctx context.Context,
	request *pb.ConnectServiceAndMirror) (*pb.Mirror, error) {
	base, base_err := models.FindMirrorById(request.MirrorId)
	if base_err != nil {
		return nil, errors.New("[Mirror] Record not found.")
	}
	_, endpoint_err := models.FindMirrorById(request.EndpointId)
	if endpoint_err != nil {
		return nil, errors.New("[Endpoint] Record not found.")
	}
	models.Connection().Model(&base).Update(
		"service_endpoint_id",
		request.EndpointId,
	)
	return base.ToProto(), nil
}

func (m *MirrorServiceServer) FindById(ctx context.Context,
	request *pb.Mirror) (*pb.Mirror, error) {
	base, base_err := models.FindMirrorById(request.Id)
	if base_err != nil {
		return nil, errors.New("[Mirror] Record not found.")
	}
	return base.ToProto(), nil
}

func (m *MirrorServiceServer) UpdateById(ctx context.Context,
	request *pb.Mirror) (*pb.Mirror, error) {
	base, base_err := models.FindMirrorById(request.Id)
	if base_err != nil {
		return nil, errors.New("[Mirror] Record not found.")
	}
	models.Connection().Model(&base).Updates(
		map[string]interface{}{
			"ipv4": request.Ipv4,
			"ipv6": request.Ipv6,
		},
	)
	return base.ToProto(), nil
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
