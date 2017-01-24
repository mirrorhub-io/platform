package controllers

import (
	models "../models"
	pb "./proto"
	log "github.com/Sirupsen/logrus"
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
	x := models.Mirror{
		Name: mirror.Name,
		IPv4: mirror.Ipv4,
		IPv6: mirror.Ipv6,
	}
	models.Connection().Create(&x)
	return x.ToProto(), nil
}
