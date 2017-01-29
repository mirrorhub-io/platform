package controllers

import (
	"errors"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"github.com/mirrorhub-io/platform/models"
	"golang.org/x/net/context"
)

type ContactServiceServer struct {
}

func (m *ContactServiceServer) Create(ctx context.Context,
	c *pb.Contact) (*pb.CreateContactResponse, error) {
	if len(c.Email) == 0 {
		return nil, errors.New("Missing email.")
	}
	if len(c.Password) < 8 {
		return nil, errors.New("Password missing or to short. (length < 8)")
	}
	contact, token := models.CreateContact(c.Name, c.Email, c.Password)
	if contact == nil {
		return nil, errors.New("Creating contact failed.")
	}
	return &pb.CreateContactResponse{
		Contact: contact.ToProto(),
		Token:   token,
	}, nil
}

func (m *ContactServiceServer) Authorize(ctx context.Context,
	request *pb.AuthorizeContactRequest) (*pb.AuthorizeContactResponse, error) {
	if len(request.Email) == 0 {
		return nil, errors.New("Missing email.")
	}
	if len(request.Password) == 0 {
		return nil, errors.New("Password missing.")
	}
	contact, token := models.AuthContactWithPassword(request.Email, request.Password)
	if contact == nil {
		return nil, errors.New("Authorize failed.")
	}
	return &pb.AuthorizeContactResponse{
		Email: contact.EMail,
		Token: token,
	}, nil
}
