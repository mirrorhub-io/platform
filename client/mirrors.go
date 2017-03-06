package client

import (
	pb "github.com/mirrorhub-io/platform/controllers/proto"
)

type MirrorClient struct {
	Client *Client
	conn   pb.MirrorServiceClient
}

func (c *MirrorClient) connection() pb.MirrorServiceClient {
	if c.conn == nil {
		c.conn = pb.NewMirrorServiceClient(c.Client.Connection())
	}
	return c.conn
}

func (c *MirrorClient) List() (*pb.MirrorGetResponse, error) {
	return c.connection().Get(
		c.Client.Context,
		&pb.ListRequest{},
	)
}

func (c *MirrorClient) FindById(id int32) (*pb.Mirror, error) {
	return c.connection().FindById(
		c.Client.Context,
		&pb.Mirror{
			Id: id,
		},
	)
}

func (c *MirrorClient) UpdateById(id int32, m *pb.Mirror) (*pb.Mirror, error) {
	m.Id = id
	return c.connection().UpdateById(
		c.Client.Context,
		m,
	)
}

func (c *MirrorClient) Connect(id int32, service_id int32) (*pb.Mirror, error) {
	return c.connection().Connect(
		c.Client.Context,
		&pb.ConnectServiceAndMirror{
			EndpointId: service_id,
			MirrorId:   id,
		},
	)
}

func (c *MirrorClient) Create(m *pb.Mirror) (*pb.Mirror, error) {
	return c.connection().Create(
		c.Client.Context,
		m,
	)
}
