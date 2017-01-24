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

type MirrorCollection struct {
	Mirrors []*Mirror
}

func MirrorList(limit int, offset int) *MirrorCollection {
	mirrors := make([]*Mirror, 0)
	Connection().Find(&mirrors)
	return &MirrorCollection{Mirrors: mirrors}
}

func (mc *MirrorCollection) ToProto() []*pb.Mirror {
	mirrors := make([]*pb.Mirror, len(mc.Mirrors))
	for i, mirror := range mc.Mirrors {
		mirrors[i] = mirror.ToProto()
	}
	return mirrors
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
