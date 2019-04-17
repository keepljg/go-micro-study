package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
	"micro-demo/api/search-api/client"
	searchSrcPb"micro-demo/srv/search/proto/search"
	fileinfoPb "micro-demo/srv/fileinfo-srv/proto/fileinfo"
	"strconv"
)

type SearchEr struct {
}


func (s *SearchEr) SearchList(ctx context.Context, req *go_api.Request, resp *go_api.Response) error{
	searchClient, ok := client.SearchFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.search-api.search.call", "search client not found")
	}
	fileInfoClient, ok := client.FileInfoFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.fileinfo-api.fileinfo.call", "fileinfo client not found")
	}

	var searchBody searchSrcPb.Request
	err := json.Unmarshal([]byte(req.Body), &searchBody)
	if err  != nil {
		return errors.InternalServerError("go.micro.api.search-api.search.call", "json Unmarshal err")
	}
	searchResp , err := searchClient.SearchList(ctx, &searchBody)
	if err != nil {
		return errors.InternalServerError("go.micro.api.search-api.search.call", err.Error())
	}
	entityIds := make([]*fileinfoPb.EntityId, 0, 10)
	for _, entityId := range searchResp.EntityIds {
		entityIds = append(entityIds, &fileinfoPb.EntityId{EntityId:strconv.Itoa(int(entityId))})
	}

	fileInfoResp, err := fileInfoClient.GetAppInfoByEntityId(ctx, &fileinfoPb.Request{
		EntityIds:entityIds,
	})

	if err != nil {
		return errors.InternalServerError("go.micro.api.fileinfo-api.fileinfo.call", err.Error())
	}
	respBody, err := json.Marshal(fileInfoResp)
	if err != nil {
		return errors.InternalServerError("go.micro.api.fileinfo-api.fileinfo.call", err.Error())
	}
	resp.StatusCode = 200
	resp.Body = string(respBody)

	return nil
}
