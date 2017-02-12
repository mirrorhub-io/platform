package client

import (
	pb "github.com/mirrorhub-io/platform/controllers/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	ContactEmail    string
	ContactPassword string
	ContactToken    string
}
