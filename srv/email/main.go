package main

import (
	"context"
	"fmt"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/opentracing/opentracing-go"
	"micro-demo/srv/email/handler"
	pb "micro-demo/srv/email/proto/email"
	"micro-demo/srv/email/subscriber"
	"time"
	tracer "github.com/hb-go/micro/pkg/opentracing"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing"
)


// 实现server.HandlerWrapper接口
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	fmt.Println("init log wrapper")
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[%v] server request: %s", time.Now(), req.Endpoint())
		return fn(ctx, req, rsp)
	}
}


func main() {
	// New Service

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"192.168.8.154:2379"}
	})
	t, closer, err := tracer.NewJaegerTracer("email-srv", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(t)
	service := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
		micro.WrapHandler(logWrapper),
		micro.Registry(reg),
		micro.WrapClient(opentracing2.NewClientWrapper()),
		micro.WrapHandler(opentracing2.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	// Register Handler
	pb.RegisterEmailHandler(service.Server(), new(handler.Email))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.email.topic", service.Server(), subscriber.SendEmail)


	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.email", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
