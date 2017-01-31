package models

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	redis "gopkg.in/redis.v5"
	"os"
	"time"
)

var redisconn *redis.Client

type Contact struct {
	gorm.Model

	Name           string
	EMail          string `gorm:"not null;unique"`
	PasswordDigest string
	Admin          bool
}

func r() *redis.Client {
	addr := "redis:6379"
	if os.Getenv("REDIS_ADDR") != "" {
		addr = os.Getenv("REDIS_ADDR")
	}
	if redisconn == nil {
		return redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
			PoolSize: 100,
		})
	}
	return redisconn
}

func (c *Contact) DelToken(token string) {
	r().Del("contact_token::" + token)
}

func (c *Contact) GenerateToken() string {
	token := uuid.NewV4().String()
	err := r().Set("contact_token::"+token, c.EMail, time.Hour*24).Err()
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
	email, err := r().Get("contact_token::" + token).Result()
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
		Id:    int32(c.ID),
		Name:  c.Name,
		Email: c.EMail,
		Admin: c.Admin,
	}
}

func CreateContact(name, email, password string) (*Contact, string) {
	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	c := &Contact{
		Name:           name,
		EMail:          email,
		PasswordDigest: string(pw),
		Admin:          false,
	}
	Connection().Create(&c)
	if Connection().NewRecord(c) {
		return nil, ""
	}
	token := c.GenerateToken()
	return c, token
}
