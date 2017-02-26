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
