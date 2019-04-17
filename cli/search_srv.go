package main

import (
	"context"
	"fmt"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	searchpb "micro-demo/srv/search/proto/search"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"192.168.1.101:2379"}
	})
	service := micro.NewService(
		micro.Registry(reg),
		)

	cli := searchpb.NewSearchService("go.micro.srv.search", service.Client())
	resp, err  := cli.SearchList(context.TODO(), &searchpb.Request{
		QueryText:"tutuapp",
		PageSize:10,
		Page:1,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(resp)
}
