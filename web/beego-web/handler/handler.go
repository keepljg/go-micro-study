package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gpmgo/gopm/modules/log"
	"strconv"

	fileinfoPb "micro-demo/srv/fileinfo-srv/proto/fileinfo"
	searchSrcPb "micro-demo/srv/search/proto/search"
)



type SearchController struct {
	beego.Controller
}


func (this *SearchController) SearchList() {
	var searchBody searchSrcPb.Request
	//searchBody.QueryText = "123"
	//searchBody.Page = 1
	//searchBody.PageSize = 10
	fmt.Println("data is " + string(this.Ctx.Input.RequestBody))
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &searchBody); err != nil {
		log.Error(err.Error())
		return
	}

	searchResp , err := ClientSearch.SearchList(context.TODO(), &searchBody)
	if err != nil {
		this.ServeJSON()
		return
	}
	entityIds := make([]*fileinfoPb.EntityId, 0, 10)
	for _, entityId := range searchResp.EntityIds {
		entityIds = append(entityIds, &fileinfoPb.EntityId{EntityId:strconv.Itoa(int(entityId))})
	}

	fileInfoResp, err := ClientFileInfo.GetAppInfoByEntityId(context.TODO(), &fileinfoPb.Request{
		EntityIds:entityIds,
	})
	this.Data["json"] = fileInfoResp
	this.ServeJSON()
}


//func ExampleCall(w http.ResponseWriter, r *http.Request) {
//	// decode the incoming request as json
//	var request map[string]interface{}
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	// call the backend service
//	exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
//	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
//		Name: request["name"].(string),
//	})
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	// we want to augment the response
//	response := map[string]interface{}{
//		"msg": rsp.Msg,
//		"ref": time.Now().UnixNano(),
//	}
//
//	// encode and write the response as json
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//}
