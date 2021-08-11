package servers

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	d "github.com/zhangshanwen/plug/server/docker"
	"log"

	pb "github.com/zhangshanwen/plug/proto"
)

type dockerServer struct {
	pb.UnimplementedDockerServer
}

func (s *dockerServer) Info(ctx context.Context, in *pb.InfoRequest) (reply *pb.InfoReply, err error) {
	res, err := d.Docker.Info(ctx)
	if err != nil {
		log.Printf("get docker info err:%v", err)
		return
	}
	reply = &pb.InfoReply{}
	b, err := json.Marshal(&res)
	if err != nil {
		log.Printf("json marshal  err:%v", err)
		return
	}
	reply.Info = b
	return
}

func (s *dockerServer) Ping(ctx context.Context, in *pb.PingRequest) (reply *pb.PingReply, err error) {
	res, err := d.Docker.Ping(ctx)
	if err != nil {
		log.Printf("get docker ping err:%v", err)
		return
	}
	reply = &pb.PingReply{}
	if err = copier.Copy(&reply, res); err != nil {
		log.Printf("copy docker ping err:%v", err)
	}
	return
}
