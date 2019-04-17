package handler

import (
	"context"
	"fmt"
	pb "micro-demo/srv/email/proto/email"
)

type Email struct {

}


func (this *Email) SendEmail(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	fmt.Println("send email by handler" +  req.Message.Message)
	//resp = &pb.Response{}
	resp.Success = 1
	return nil
}