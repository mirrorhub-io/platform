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
	request *pb.ListRequest) (*pb.MirrorGetResponse, error) {
	if _, err := contactAuth(ctx); err != nil {
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
	if request.MirrorId == request.EndpointId {
		return nil, errors.New("MirrorId shouldn't equal with EndpointId")
	}
	base, base_err := models.FindMirrorById(request.MirrorId)
	if base_err != nil {
		return nil, errors.New("[Mirror] Record not found.")
	}
	_, endpoint_err := models.FindMirrorById(request.EndpointId)
	if endpoint_err != nil {
		return nil, errors.New("[Endpoint] Record not found.")
	}
	base.ServiceEndpointID = request.EndpointId
	models.Connection().Save(&base)
	return base.ToProto(), nil
}

func (m *MirrorServiceServer) FindById(ctx context.Context,
	request *pb.Mirror) (*pb.Mirror, error) {
	if _, err := contactAuth(ctx); err != nil {
		return nil, err
	}
	base, base_err := models.FindMirrorById(request.Id)
	if base_err != nil {
		return nil, errors.New("[Mirror] Record not found.")
	}
	return base.ToProto(), nil
}

func (m *MirrorServiceServer) UpdateById(ctx context.Context,
	request *pb.Mirror) (*pb.Mirror, error) {
	contact, err := contactAuth(ctx)
	if err != nil {
		return nil, err
	}
	base, base_err := models.FindMirrorById(request.Id)
	if base_err != nil {
		return nil, errors.New("[Mirror] Record not found.")
	}
	if base.ContactID != int32(contact.ID) && contact.Admin == false {
		return nil, errors.New("Insufficient permissions")
	}
	models.Connection().Model(&base).Updates(
		models.MirrorFromProto(request),
	)
	return base.ToProto(), nil
}

func (m *MirrorServiceServer) Create(ctx context.Context, mirror *pb.Mirror) (*pb.Mirror, error) {
	contact, err := contactAuth(ctx)
	if err != nil {
		return nil, err
	}
	mirror.ContactId = int32(contact.ID)
	x := models.MirrorFromProto(mirror)
	models.Connection().Create(&x)
	return x.ToProto(), nil
}

func contactAuth(ctx context.Context) (*models.Contact, error) {
	contact, _, err := AuthContact(ctx)
	if err != nil {
		return contact, err
	}
	if contact == nil {
		return contact, errors.New("Unauthorized")
	}
	return contact, nil
}
