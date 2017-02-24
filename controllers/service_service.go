package controllers

import (
	"errors"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"github.com/mirrorhub-io/platform/models"
	"golang.org/x/net/context"
	"strings"
)

type ServiceServiceServer struct {
}

func (m *ServiceServiceServer) Get(ctx context.Context,
	request *pb.ListRequest) (*pb.MirrorGetResponse, error) {
	if _, err := contactAuth(ctx); err != nil {
		return nil, err
	}
	services := make([]*pb.Service, 0)
	services = models.ServiceList(10, 0).ToProto()
	return &pb.ServiceGetResponse{
		Services: services,
	}, nil
}

func (m *ServiceServiceServer) Find(ctx context.Context,
	request *pb.Service) (*pb.Service, error) {
	if _, err := contactAuth(ctx); err != nil {
		return nil, err
	}
	base, base_err := models.FindServiceById(request.Id)
	if base_err != nil {
		return nil, errors.New("[Service] Record not found.")
	}
	return base.ToProto(), nil
}

func (m *ServiceServiceServer) Update(ctx context.Context,
	request *pb.Service) (*pb.Service, error) {
	contact, err := contactAuth(ctx)
	if err != nil {
		return nil, err
	}
	base, base_err := models.FindServiceById(request.Id)
	if base_err != nil {
		return nil, errors.New("[Service] Record not found.")
	}
	if contact.Admin == false {
		return nil, errors.New("Insufficient permissions")
	}
	models.Connection().Model(&base).Updates(
		&models.Service{
			Name:     request.Name,
			Storage:  request.Storage,
			FileList: strings.Join(request.FileList, ","),
		},
	)
	return base.ToProto(), nil
}

func (m *ServiceServiceServer) Create(ctx context.Context, mirror *pb.Service) (*pb.Service, error) {
	contact, err := contactAuth(ctx)
	if err != nil {
		return nil, err
	}
	x = &models.Service{
		Name:     request.Name,
		Storage:  request.Storage,
		FileList: strings.Join(request.FileList, ","),
	}
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
