package models

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	utils "github.com/mirrorhub-io/platform/utils"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Contact struct {
	gorm.Model

	Name           string
	EMail          string `gorm:"not null;unique"`
	PasswordDigest string
	Admin          bool
}

func (c *Contact) DelToken(token string) {
	utils.Redis().Del("contact_token::" + token)
}

func (c *Contact) GenerateToken() string {
	token := uuid.NewV4().String()
	err := utils.Redis().Set("contact_token::"+token, c.EMail, time.Hour*24).Err()
	if err != nil {
		log.Error(err)
	}
	return token
}

func FindContactByEmail(email string) (*Contact, error) {
	c := &Contact{EMail: email}
	Connection().Where(&c).First(&c)
	if Connection().NewRecord(c) {
		return nil, errors.New("Record not present.")
	}
	return c, nil
}

func (c *Contact) Update(cpb *pb.Contact, token string) (*Contact, string) {
	if cpb.Email != "" {
		c.EMail = cpb.Email
	}
	if cpb.Name != "" {
		c.Name = cpb.Name
	}
	if len(cpb.Password) > 0 {
		pw, _ := bcrypt.GenerateFromPassword([]byte(cpb.Password), 10)
		c.PasswordDigest = string(pw)
	}
	Connection().Save(c)
	c.DelToken(token)
	return c, c.GenerateToken()
}

func AuthContactWithToken(token string) (*Contact, error) {
	email, err := utils.Redis().Get("contact_token::" + token).Result()
	if err != nil {
		return nil, errors.New("Token not present.")
	}
	return FindContactByEmail(email)
}

func AuthContactWithPassword(email, password string) (*Contact, string, error) {
	c, err := FindContactByEmail(email)
	if err != nil {
		return c, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(c.PasswordDigest), []byte(password))
	if err != nil {
		return nil, "", errors.New("Password error.")
	}
	val := c.GenerateToken()
	return c, val, nil
}

func (c *Contact) ToProto() *pb.Contact {
	return &pb.Contact{
		Id:      int32(c.ID),
		Name:    c.Name,
		Email:   c.EMail,
		Admin:   c.Admin,
		Mirrors: c.Mirrors().ToProto(),
	}
}

func (c *Contact) Mirrors() *MirrorCollection {
	m := &Mirror{ContactID: int32(c.ID)}
	mirrors := make([]*Mirror, 0)
	Connection().Where(&m).Find(&mirrors)
	return &MirrorCollection{
		Mirrors: mirrors,
	}
}

func CreateContact(name, email, password string) (*Contact, string, error) {
	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	c := &Contact{
		Name:           name,
		EMail:          email,
		PasswordDigest: string(pw),
		Admin:          false,
	}
	err := Connection().Create(&c).GetErrors()
	if len(err) > 0 {
		return nil, "", joinErrors(err)
	}
	if Connection().NewRecord(c) {
		return nil, "", errors.New("Contact create error.")
	}
	token := c.GenerateToken()
	return c, token, nil
}
