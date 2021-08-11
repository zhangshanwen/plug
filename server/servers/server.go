package servers

import (
	"google.golang.org/grpc"

	pb "github.com/zhangshanwen/plug/proto"
)

func InitServers(s *grpc.Server) {
	pb.RegisterImageServer(s, &imagesServer{})
	pb.RegisterDockerServer(s, &dockerServer{})
	pb.RegisterImageAliveServer(s, &imagesAliveServer{})

}
