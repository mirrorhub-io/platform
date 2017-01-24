package models

import (
	pb "../controllers/proto"
	"github.com/jinzhu/gorm"
	"time"
)

type Mirror struct {
	gorm.Model

	ContactID int32
	IPv4      string
	IPv6      string
	Domain    string
	Location  string
	Name      string

	Traffic          int64
	TrafficResetDay  int32
	Bandwidth        int32
	AvailableStorage int64

	ClientToken string

	Services []Service `gorm:"many2many:services_mirrors;"`
}

func (m *Mirror) ToProto() *pb.Mirror {
	return &pb.Mirror{
		Name:        m.Name,
		Ipv4:        m.IPv4,
		Ipv6:        m.IPv6,
		CreatedAt:   m.CreatedAt.Unix(),
		OnlineSince: time.Now().Unix(),
	}
}

type Service struct {
	gorm.Model

	Name               string
	Storage            int64
	TrafficConsumption int64
}

type Contact struct {
	gorm.Model

	Name    string
	EMail   string
	Mirrors []Mirror
}
