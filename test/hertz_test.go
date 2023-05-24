package test

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"testing"
)

func TestFileUploadHertz(t *testing.T) {
	client, err := client.NewClient()
	if err != nil {
		return
	}
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodPost)
	req.SetRequestURI("http://127.0.0.1:8080/singleFile")
	req.SetFile("file", "./client/upload_file/file1.txt")

	err = client.Do(context.Background(), req, res)
	if err != nil {
		return
	}
	fmt.Println(err, string(res.Body()))
}
