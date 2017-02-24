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
	request *pb.ServiceGetRequest) (*pb.ServiceGetResponse, error) {
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
			FileList: strings.Join(request.Files, ","),
		},
	)
	return base.ToProto(), nil
}

func (m *ServiceServiceServer) Create(ctx context.Context, request *pb.Service) (*pb.Service, error) {
	contact, err := contactAuth(ctx)
	if err != nil {
		return nil, err
	}
	if contact.Admin == false {
		return nil, errors.New("Insufficient permissions")
	}
	x := &models.Service{
		Name:     request.Name,
		Storage:  request.Storage,
		FileList: strings.Join(request.Files, ","),
	}
	models.Connection().Create(&x)
	return x.ToProto(), nil
}
