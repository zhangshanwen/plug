package client

import (
	"context"
	"io"
	"log"
	"testing"

	pb "github.com/zhangshanwen/plug/proto"
)

func TestGetImages(t *testing.T) {
	ctx := context.Background()
	c, err := New(ctx, defaultTimeOut, defaultAddress)
	if err != nil {
		t.Fatalf("new client err:%v", err)
	}
	defer c.Close()
	pbc := pb.NewImageClient(c.Conn)
	images, err := pbc.ImageList(ctx, &pb.ImageListRequest{All: true})
	if err != nil {
		t.Fatalf("get images  err:%v", err)
	}
	t.Log(images)
}

func TestInfo(t *testing.T) {
	ctx := context.Background()
	c, err := New(ctx, defaultTimeOut, defaultAddress)
	if err != nil {
		t.Fatalf("new client err:%v", err)
	}
	defer c.Close()
	pbc := pb.NewDockerClient(c.Conn)
	images, err := pbc.Info(ctx, &pb.InfoRequest{})
	if err != nil {
		t.Fatalf("get info err:%v", err)
	}
	t.Log(images)
	ping, err := pbc.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		t.Fatalf("get info err:%v", err)
	}
	t.Log(ping)
}

func TestImagePull(t *testing.T) {
	ctx := context.Background()
	c, err := NewKeep(ctx, defaultTimeOut, defaultAddress)
	if err != nil {
		t.Fatalf("new client err:%v", err)
	}
	defer c.Close()
	pbc := pb.NewImageAliveClient(c.Conn)
	stream, err := pbc.ImagePull(ctx, &pb.ImagePullRequest{RefStr: "mysql"})
	if err != nil {
		log.Fatalf("ImagePullRequest err  %v", err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF || err != nil {
			break
		}
		t.Log(string(feature.Body), "-----------")
	}
}
