package controllers

import (
	models "../models"
	pb "./proto"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type mirrorServiceServer struct {
}

func (m *mirrorServiceServer) Get(ctx context.Context, request *pb.MirrorGetRequest) (*pb.MirrorGetResponse, error) {
	log.Info("test")
	return &pb.MirrorGetResponse{
		Mirrors: []*pb.Mirror{
			&pb.Mirror{
				Name:        "Moo",
				OnlineSince: time.Now().Unix(),
			},
		},
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
