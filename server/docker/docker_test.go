package docker

import (
	"testing"
)

func TestDocker(t *testing.T) {
	var err error
	if err = InitDocker(); err != nil {
		t.Fatalf("init docker client err:%v", err)
	}
	defer Docker.Close()
	//t.Log(Docker.ImageList(ctx, types.ImageListOptions{All: true}))
	//t.Log(Docker.ImageHistory(ctx, "sha256:0cc4e230376eac5417da8c6a7da73ee53494bbc82b4143cb9ce570cfc2ec2b29"))
	//t.Log(Docker.ImageSearch(ctx, "redis", types.ImageSearchOptions{Limit: 100}))
	//t.Log(Docker.ImageSearch(context.Background(), "mysql", types.ImageSearchOptions{Limit: 100}))
	//reader, err := Docker.ImagePull(context.Background(), "mysql", types.ImagePullOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//io.Copy(os.Stdout, reader)a
	Docker.ContainerList()

}
