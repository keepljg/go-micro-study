package handler

import (
	"github.com/micro/go-micro/client"
	filePb "micro-demo/srv/fileinfo-srv/proto/fileinfo"
	searchPb "micro-demo/srv/search/proto/search"
)

var (
	ClientSearch searchPb.SearchService
	ClientFileInfo filePb.FileInfoService
)

func InitClient() {
	ClientSearch = searchPb.NewSearchService("go.micro.srv.search", client.DefaultClient)
	ClientFileInfo = filePb.NewFileInfoService("go.micro.srv.fileinfo", client.DefaultClient)
}