package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"protos"
)

type QslQuery struct {

}

func main()  {

	certFile := ""
	creds, _ := credentials.NewClientTLSFromFile(certFile, "")


	conn, err := grpc.Dial(
		"172.16.16.250:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("连接创建失败", err)
	}

	c := protos.NewQSLEngineClient(conn)
	queryRequest := &protos.SearchRequest{
		AppKey: "insight_r2_backend",
		AppUser: "1",
		Qsl: "show alias",
		Timeout: 300,
		QueryType: protos.SearchRequest_ADVANCED,
	}
	//statusRequest := &protos.StatusRequest{JobId: "d69058460a5711ecbf21fa163ee6c9a7"}
	r, err := c.Query(context.Background(), queryRequest)
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}

	fmt.Printf("调用成功: %s", r.Status)
}
