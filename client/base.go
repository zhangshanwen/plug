package client

import (
	"context"
	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc"
	"time"
)

type Client struct {
	cel  context.CancelFunc
	Ctx  context.Context
	Conn *grpc.ClientConn
}

const (
	defaultAddress = "localhost:50051"
	defaultTimeOut = time.Second * 3
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func New(ctx context.Context, dialTimeOut time.Duration, address string) (c Client, err error) {
	c.Ctx, c.cel = context.WithTimeout(ctx, dialTimeOut)
	c.Conn, err = grpc.DialContext(c.Ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	return
}
func NewKeep(ctx context.Context, dialTimeOut time.Duration, address string) (c Client, err error) {
	c.Ctx, c.cel = context.WithTimeout(ctx, dialTimeOut)
	c.Conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		return
	}
	return

}

func (c *Client) Default() (err error) {
	c.Ctx, c.cel = context.WithTimeout(context.Background(), defaultTimeOut)
	c.Conn, err = grpc.DialContext(c.Ctx, defaultAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	return
}

func (c *Client) Close() {
	defer c.cel()
	defer c.Conn.Close()

}
