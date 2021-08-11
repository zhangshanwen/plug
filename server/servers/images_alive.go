package servers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/jinzhu/copier"
	pb "github.com/zhangshanwen/plug/proto"
	d "github.com/zhangshanwen/plug/server/docker"
	"io"
	"log"
)

type imagesAliveServer struct {
	pb.UnimplementedImageAliveServer
}

func (s *imagesAliveServer) ImagePull(in *pb.ImagePullRequest, steam pb.ImageAlive_ImagePullServer) (err error) {
	option := types.ImagePullOptions{}
	if err = copier.Copy(&option, &in); err != nil {
		return
	}
	res, err := d.Docker.ImagePull(context.Background(), in.RefStr, option)
	if err != nil {
		log.Printf("get images list err:%v", err)
		return
	}
	reply := &pb.ImagePullReply{}
	size := 1024
	reply.Body = make([]byte, size)
	for {
		length, err := res.Read(reply.Body)
		if err == io.EOF || err != nil {
			fmt.Println("length=", length, "err", err)
			res.Close()
			break
		}
		if length <= 0 {
			continue
		}
		steam.Send(reply)
	}
	return
}
