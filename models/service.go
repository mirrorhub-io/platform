package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"strings"
)

type Service struct {
	gorm.Model

	Name     string
	Storage  int64
	FileList string
}

type ServiceCollection struct {
	Services []*Service
	Proto    []*pb.Service
}

func ServiceFromProto(s *pb.Service) *Service {
	return &Service{
		Name:    s.Name,
		Storage: s.Storage,
	}
}

func FindServiceById(id int32) (*Service, error) {
	se := &Service{}
	if Connection().Where(
		"id = ?",
		id,
	).First(&se).RecordNotFound() {
		return nil, errors.New("Record not found.")
	}
	return se, nil
}

func ServiceList(limit int, offset int) *ServiceCollection {
	services := make([]*Service, 0)
	Connection().Find(&services)
	return &ServiceCollection{Services: services}
}

func (sc *ServiceCollection) ToServices() []*Service {
	services := make([]*Service, len(sc.Proto))
	for i, service := range sc.Proto {
		services[i] = ServiceFromProto(service)
	}
	return services
}

func (sc *ServiceCollection) ToProto() []*pb.Service {
	services := make([]*pb.Service, len(sc.Services))
	for i, service := range sc.Services {
		services[i] = service.ToProto()
	}
	return services
}

func (s *Service) Files() []string {
	return strings.Split(s.FileList, ",")
}

func (s *Service) ToProto() *pb.Service {
	return &pb.Service{
		Id:      int32(s.ID),
		Name:    s.Name,
		Storage: s.Storage,
		Files:   s.Files(),
	}
}
