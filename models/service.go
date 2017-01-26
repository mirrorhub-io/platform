package models

import (
	"github.com/jinzhu/gorm"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
)

type Service struct {
	gorm.Model

	Name               string
	Storage            int64
	TrafficConsumption int64
}

type ServiceCollection struct {
	Services []*Service
	Proto    []*pb.Service
}

func ServiceFromProto(s *pb.Service) *Service {
	return &Service{
		Name:               s.Name,
		Storage:            s.Storage,
		TrafficConsumption: s.TrafficConsumption,
	}
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

func (s *Service) ToProto() *pb.Service {
	return &pb.Service{
		Id:                 int32(s.ID),
		Name:               s.Name,
		Storage:            s.Storage,
		TrafficConsumption: s.TrafficConsumption,
	}
}
