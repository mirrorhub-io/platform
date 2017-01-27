package models

import (
	"github.com/jinzhu/gorm"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
)

type Mirror struct {
	gorm.Model

	ContactID        int32
	IPv4             string `gorm:"not null"`
	IPv6             string `gorm:"not null"`
	Domain           string `gorm:"not null"`
	LocalDestination string `gorm:"not null"`
	Name             string `gorm:"not null"`

	Traffic   int64
	Bandwidth int64
	Storage   int64

	ClientToken string `gorm:"not null;unique"`

	ServiceEnpointID int32
	ServiceID        int32
}

type MirrorCollection struct {
	Mirrors []*Mirror
}

func MirrorList(limit int, offset int) *MirrorCollection {
	mirrors := make([]*Mirror, 0)
	Connection().Find(&mirrors)
	return &MirrorCollection{Mirrors: mirrors}
}

func MirrorFromProto(p *pb.Mirror) *Mirror {
	return &Mirror{
		Name:        p.Name,
		IPv4:        p.Ipv4,
		IPv6:        p.Ipv6,
		Domain:      p.Domain,
		ContactID:   p.ContactId,
		Traffic:     p.Traffic,
		ClientToken: p.ClientToken,
		Bandwidth:   p.Bandwidth,
	}
}

func (mc *MirrorCollection) ToProto() []*pb.Mirror {
	mirrors := make([]*pb.Mirror, len(mc.Mirrors))
	for i, mirror := range mc.Mirrors {
		mirrors[i] = mirror.ToProto()
	}
	return mirrors
}

func (m *Mirror) FetchServices() *ServiceCollection {
	services := make([]*Service, 0)
	Connection().Model(&m).Related(&services, "Services")
	return &ServiceCollection{
		Services: services,
	}
}

func (m *Mirror) ToProto() *pb.Mirror {
	return &pb.Mirror{
		Name:        m.Name,
		Ipv4:        m.IPv4,
		Ipv6:        m.IPv6,
		Services:    m.FetchServices().ToProto(),
		Domain:      m.Domain,
		ContactId:   m.ContactID,
		Traffic:     m.Traffic,
		ClientToken: m.ClientToken,
		Bandwidth:   m.Bandwidth,
		CreatedAt:   m.CreatedAt.Unix(),
	}
}
