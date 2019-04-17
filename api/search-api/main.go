package main

import (
	tracer "github.com/hb-go/micro/pkg/opentracing"
	"github.com/micro/go-api"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"log"
	"micro-demo/api/search-api/client"
	"micro-demo/api/search-api/handler"
	pb "micro-demo/api/search-api/proto/search"
)

func main() {
	// New Service
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"192.168.8.154:2379"}
	})

	t, closer, err := tracer.NewJaegerTracer("search-api", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(t)
	service := micro.NewService(
		micro.Name("go.micro.api.search"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.WrapClient(opentracing2.NewClientWrapper()),
		micro.WrapHandler(opentracing2.NewHandlerWrapper()),

	)
	service.Server().Init()
	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.SearchWrapper(service)),
		//micro.WrapHandler(ServerWrapper),
	)

	// Register Handler
	pb.RegisterSearchErHandler(service.Server(), new(handler.SearchEr), api.WithEndpoint(&api.Endpoint{
		// The RPC method
		Name:        "Search.SearchEr.SearchList",
		Description: "",
		Handler:     "api",
		Host:        nil,
		Method:      []string{"post"},
		Path:        []string{"/searchEr/searchList"},
	}))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}


