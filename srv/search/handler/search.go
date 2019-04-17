package handler

import (
	"context"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/micro/go-micro"
	"math/rand"
	pb "micro-demo/srv/search/proto/search"
	"time"
	emailPb "micro-demo/srv/email/proto/email"
)

type SearchEr struct {
	micro.Publisher
}


// start 到 end 的随机数
func RangeInt(start, end int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(end -start + 1) + start
}

func (s *SearchEr) SearchList(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	queryText := req.QueryText
	log.Info(queryText)
	entityIds := make([]int64, 0)
	for i := 0; i< int(req.PageSize); i ++ {
		entityIds = append(entityIds, RangeInt(10, 100000))
	}
	if err := s.Publish(ctx, &emailPb.Message{Message:"this is a email"}); err != nil {
		fmt.Println(err)
	}
	resp.EntityIds = entityIds
	resp.Page = req.Page
	resp.CountAll = req.Page * req.PageSize
	return nil
}
