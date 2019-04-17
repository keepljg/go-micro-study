package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	emailPb "micro-demo/srv/email/proto/email"
	"time"
)


type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[email] client request to service: %s method: %s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

// 实现client.Wrapper，充当日志包装器
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}


func main()  {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"192.168.8.154:2379"}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.WrapClient(logWrap),
	)
	service.Init()
	cli := emailPb.NewEmailService("go.micro.srv.email", service.Client())
	resp, err := cli.SendEmail(context.TODO(), &emailPb.Request{Message:&emailPb.Message{Message:"123"}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	publisher := micro.NewPublisher("go.micro.srv.email.topic", service.Client())

	err = publisher.Publish(context.TODO(), &emailPb.Message{Message:"this is a email"})

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(1 * time.Second)

}
