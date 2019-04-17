package main

import (
	tracer "github.com/hb-go/micro/pkg/opentracing"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/opentracing/opentracing-go"
	"micro-demo/srv/fileinfo-srv/handler"
	pb "micro-demo/srv/fileinfo-srv/proto/fileinfo"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

func main() {

	t, closer, err := tracer.NewJaegerTracer("fileinfo-srv", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(t)
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"192.168.8.154:2379"}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.fileinfo"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.WrapClient(opentracing2.NewClientWrapper()),
		micro.WrapHandler(opentracing2.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	// Register Handler
	pb.RegisterFileInfoHandler(service.Server(), new(handler.GetFileInfoEr))

	//// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.tutu.srv.fileinfo-srv", service.Server(), new(subscriber.Example))
	//
	//// Register Function as Subscriber
	//micro.RegisterSubscriber("go.tutu.srv.fileinfo-srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
