package main

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-web"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"google.golang.org/grpc/metadata"
	"log"
	"micro-demo/web/beego-web/handler"
	"time"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"192.168.8.154:2379"}
	})
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.beego-web"),
		web.Version("latest"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Registry(reg),

	)

	handler.InitClient()
	// initialise service
	if err := service.Init(
		); err != nil {
		log.Fatal(err)
	}

	router := beego.NewControllerRegister()
	router.Add("/search/searchlist", &handler.SearchController{}, "post:SearchList")

	//beeapp := beego.NewApp()
	beego.BConfig.CopyRequestBody = true

	// register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.Handle("/", router)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}


func InitTracer(zipkinURL string, hostPort string, serviceName string) {
	collector, err := zipkin.NewHTTPCollector(zipkinURL) // 创建一个 Zipkin HTTP 后端收集器
	if err != nil {
		log.Fatalf("unable to create Zipkin HTTP collector: %v", err)
		return
	}
	tracer, err := zipkin.NewTracer(   // 创建一个基于 Zipkin 收集器的记录器
		zipkin.NewRecorder(collector, false, hostPort, serviceName), // 创建一个基于 Zipkin 收集器的记录器
	)
	if err != nil {
		log.Fatalf("unable to create Zipkin tracer: %v", err)
		return
	}
	opentracing.InitGlobalTracer(tracer)
	return
}


func ServerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		operationName := req.Method()

		//extract metadata to context
		ctx = ContextFromGRPC(ctx, opentracing.GlobalTracer(), operationName)

		//get span from context metadata
		span := opentracing.SpanFromContext(ctx)
		if span == nil {
			//create new root span
			//span = opentracing.StartSpan(operationName)
			return nil
		}

		//span.SetOperationName(operationName)
		defer span.Finish()

		ext.SpanKindRPCServer.Set(span)
		span.SetTag("test tag", "fuck")

		log.Printf("[Trace Wrapper] Before serving request method: %v\n", req.Method())
		err := fn(ctx, req, rsp)
		log.Printf("[Trace Wrapper] After serving request. TraceId: %v\n", opentracing.GlobalTracer())

		return err
	}
}

func ContextFromGRPC(ctx context.Context, tracer opentracing.Tracer, operationName string) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)

	var span opentracing.Span
	wireContext, err := tracer.Extract(opentracing.TextMap, metadataReader{&md})
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		log.Printf("metadata error %s\n", err)
	}
	span = tracer.StartSpan(operationName, ext.RPCServerOption(wireContext))
	return opentracing.ContextWithSpan(ctx, span)
}


// A type that conforms to opentracing.TextMapReader and
// opentracing.TextMapWriter.
type metadataReader struct {
	*metadata.MD
}

func (w metadataReader) ForeachKey(handler func(key, val string) error) error {
	for k, vals := range *w.MD {
		for _, v := range vals {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}