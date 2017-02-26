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
	c *pb.Contact) (*pb.ContactResponse, error) {
	if len(c.Email) == 0 {
		return nil, errors.New("Missing email.")
	}
	if len(c.Password) < 8 {
		return nil, errors.New("Password missing or to short. (length min 8 chars)")
	}
	contact, token, err := models.CreateContact(c.Name, c.Email, c.Password)
	if err != nil {
		return nil, err
	}
	if contact == nil {
		return nil, errors.New("Creating contact failed.")
	}
	return newContactResponse(contact, token)
}

func (m *ContactServiceServer) Update(ctx context.Context,
	c *pb.Contact) (*pb.ContactResponse, error) {
	contact, token, err := AuthContact(ctx)
	if err != nil {
		return nil, err
	}
	new_contact, token := contact.Update(c, token)
	return newContactResponse(new_contact, token)
}

func (m *ContactServiceServer) Get(ctx context.Context,
	c *pb.Contact) (*pb.ContactResponse, error) {
	contact, token, err := AuthContact(ctx)
	if err != nil {
		return nil, err
	}
	return newContactResponse(contact, token)
}

func (m *ContactServiceServer) Authorize(ctx context.Context,
	request *pb.AuthorizeContactRequest) (*pb.ContactResponse, error) {
	if len(request.Email) == 0 {
		return nil, errors.New("Missing email.")
	}
	if len(request.Password) == 0 {
		return nil, errors.New("Password missing.")
	}
	contact, token, err := models.AuthContactWithPassword(request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	if contact == nil {
		return nil, errors.New("Authorize failed.")
	}
	return newContactResponse(contact, token)
}

func newContactResponse(contact *models.Contact, token string) (*pb.ContactResponse, error) {
	return &pb.ContactResponse{
		Contact: contact.ToProto(),
		Token:   token,
	}, nil
}
