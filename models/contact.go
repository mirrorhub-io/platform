package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	redis "gopkg.in/redis.v5"
	"os"
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
	addr := "redis"
	if os.Getenv("REDIS_ADDR") != "" {
		addr = os.Getenv("REDIS_ADDR")
	}
	if redisconn == nil {
		redisconn = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
			PoolSize: 100,
		})
	}
	return redisconn
}

func (c *Contact) GenerateToken() string {
	token := uuid.NewV4().String()
	r().Set("contact_token::"+c.EMail, token, 86400)
	return token
}

func (c *Contact) Token() (string, error) {
	return r().Get("contact_token::" + c.EMail).Result()
}

func AuthContactWithToken(email, token string) (*Contact, string) {
	c := &Contact{EMail: email}
	Connection().Where(&c).First(&c)
	if Connection().NewRecord(c) {
		return nil, ""
	}
	val, err := c.Token()
	if err != nil {
		return nil, ""
	}
	return c, val
}

func AuthContactWithPassword(email, password string) (*Contact, string) {
	c := &Contact{EMail: email}
	Connection().Where(&c).First(&c)
	if Connection().NewRecord(c) {
		return nil, ""
	}
	err := bcrypt.CompareHashAndPassword([]byte(c.PasswordDigest), []byte(password))
	if err != nil {
		return nil, ""
	}
	val := c.GenerateToken()
	return c, val
}

func CreateContact(name, email, password string) *Contact {
	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	c := &Contact{
		Name:           name,
		EMail:          email,
		PasswordDigest: string(pw),
		Admin:          false,
	}
	Connection().Create(&c)
	if Connection().NewRecord(c) {
		return nil
	}
	c.GenerateToken()
	return c
}
