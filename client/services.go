package client

import (
	pb "github.com/mirrorhub-io/platform/controllers/proto"
)

type ServiceClient struct {
	Client *Client
	conn   pb.ServiceServiceClient
}

func (c *ServiceClient) connection() pb.ServiceServiceClient {
	if c.conn == nil {
		c.conn = pb.NewServiceServiceClient(c.Client.Connection())
	}
	return c.conn
}

func (c *ServiceClient) List() (*pb.ServiceGetResponse, error) {
	return c.connection().Get(
		c.Client.Context,
		&pb.ServiceGetRequest{},
	)
}

func (c *ServiceClient) Create(s *pb.Service) (*pb.Service, error) {
	return c.connection().Create(
		c.Client.Context,
		s,
	)
}
