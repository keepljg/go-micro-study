package client

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	searchPb "micro-demo/srv/search/proto/search"
	filePb "micro-demo/srv/fileinfo-srv/proto/fileinfo"
)

type searchKey struct {}

type fileinfoKey struct {}


// FromContext retrieves the client from the Context
func SearchFromContext(ctx context.Context) (searchPb.SearchService, bool) {
	c, ok := ctx.Value(searchKey{}).(searchPb.SearchService)
	return c, ok
}


func FileInfoFromContext(ctx context.Context) (filePb.FileInfoService, bool) {
	c, ok := ctx.Value(fileinfoKey{}).(filePb.FileInfoService)
	return c, ok
}




// Client returns a wrapper for the ExampleClient
func SearchWrapper(service micro.Service) server.HandlerWrapper {
	clientSearch := searchPb.NewSearchService("go.micro.srv.search", service.Client())
	clientFileInfo := filePb.NewFileInfoService("go.micro.srv.fileinfo", service.Client())
	fmt.Println("init client")
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, searchKey{}, clientSearch)
			ctx = context.WithValue(ctx, fileinfoKey{}, clientFileInfo)
			return fn(ctx, req, rsp)
		}
	}
}
