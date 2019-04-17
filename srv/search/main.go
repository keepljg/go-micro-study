package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/opentracing/opentracing-go"
	"micro-demo/srv/search/handler"
	pb "micro-demo/srv/search/proto/search"
	tracer "github.com/hb-go/micro/pkg/opentracing"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

func main() {

	t, closer, err := tracer.NewJaegerTracer("search-srv", "127.0.0.1:6831")
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
		micro.Name("go.micro.srv.search"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.WrapClient(opentracing2.NewClientWrapper()),
		micro.WrapHandler(opentracing2.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()
	publisher := micro.NewPublisher("go.micro.srv.email.topic", service.Client())
	// Register Handler
	pb.RegisterSearchHandler(service.Server(), &handler.SearchEr{publisher})

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.tutu.srv.search", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.tutu.srv.search", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
