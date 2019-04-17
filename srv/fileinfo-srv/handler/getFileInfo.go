package handler

import (
	"context"
	"fmt"
	pb "micro-demo/srv/fileinfo-srv/proto/fileinfo"
)

type GetFileInfoEr struct {
}

func (g *GetFileInfoEr) GetAppInfoByEntityId(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	for _, v := range req.EntityIds {
		resp.AppInfos = append(resp.AppInfos, &pb.AppInfo{
			Name:                 fmt.Sprintf("tutuapp %s", v.EntityId),
			IconUrl:              fmt.Sprintf("http://tutuapp.com/icon/%s", v.EntityId),
			Screens:              []string{""},
			DownloadUrl:          fmt.Sprintf("http://tutuapp.com/download/%s", v.EntityId),
		})
	}
	return nil
}
