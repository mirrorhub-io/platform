package client

import (
	pb "github.com/mirrorhub-io/platform/controllers/proto"
)

type ContactClient struct {
	Client *Client
	conn   pb.ContactServiceClient
}

func (c *ContactClient) connection() pb.ContactServiceClient {
	if c.conn == nil {
		c.conn = pb.NewContactServiceClient(c.Client.Connection())
	}
	return c.conn
}

func (c *ContactClient) Authorize() (*pb.ContactResponse, error) {
	a, err := c.connection().Authorize(
		c.Client.Context,
		&pb.AuthorizeContactRequest{
			Email:    c.Client.ContactEmail,
			Password: c.Client.ContactPassword,
		},
	)
	if err != nil {
		return nil, err
	}
	return c.UpdateSession(a), nil
}

func (c *ContactClient) UpdateSession(a *pb.ContactResponse) *pb.ContactResponse {
	c.Client.ContactToken = a.Token
	c.Client.PrepareHeader()
	return a
}

func (c *ContactClient) Create(name, email, password string) (*pb.ContactResponse, error) {
	a, err := c.connection().Create(
		c.Client.Context,
		&pb.Contact{
			Name:     name,
			Email:    email,
			Password: password,
		},
	)
	if err != nil {
		return nil, err
	}
	return c.UpdateSession(a), nil
}
