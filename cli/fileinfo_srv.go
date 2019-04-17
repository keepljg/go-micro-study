package main

import (
	"context"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pbfileinfosrv "micro-demo/srv/fileinfo-srv/proto/fileinfo"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"192.168.8.154:2379"}
	})
	service := micro.NewService(
		micro.Registry(reg),
		)

	cli := pbfileinfosrv.NewFileInfoService("go.micro.srv.fileinfo", service.Client())
	entityIds := []*pbfileinfosrv.EntityId{&pbfileinfosrv.EntityId{EntityId:"123"}, &pbfileinfosrv.EntityId{EntityId:"234"}}
	resp, err  := cli.GetAppInfoByEntityId(context.TODO(), &pbfileinfosrv.Request{EntityIds:entityIds})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(resp)
}
