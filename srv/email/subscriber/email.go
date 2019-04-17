package subscriber

import (
	"context"
	"fmt"
	emailPb "micro-demo/srv/email/proto/email"
)

//type Email struct{}


//func (this *Email) Handle(ctx context.Context, msg *pb.Message) error {
//	fmt.Println("send email %s", msg.Message)
//	return nil
//}



func SendEmail(ctx context.Context, message *emailPb.Message) error {
	fmt.Printf("Send email by broker %+v\n", message.Message)
	return nil
}