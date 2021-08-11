package servers

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/jinzhu/copier"

	pb "github.com/zhangshanwen/plug/proto"
	d "github.com/zhangshanwen/plug/server/docker"
)

type imagesServer struct {
	pb.UnimplementedImageServer
}

func (s *imagesServer) ImageList(ctx context.Context, in *pb.ImageListRequest) (reply *pb.ImageListReply, err error) {
	args := filters.Args{}
	for _, v := range in.Filters {
		args.Add(v.Key, v.Value)
	}
	option := types.ImageListOptions{All: in.All, Filters: args}
	res, err := d.Docker.ImageList(ctx, option)
	if err != nil {
		log.Printf("get images list err:%v", err)
		return
	}
	reply = &pb.ImageListReply{}
	if err = copier.Copy(&reply.List, &res); err != nil {
		log.Printf("copy images list err:%v", err)
		return
	}
	return
}
