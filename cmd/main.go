package main

import (
	"fmt"
	"os"

	api "github.com/siangyeh8818/golang.exporter.alibabacloud/internal/api"
	"github.com/siangyeh8818/golang.exporter.alibabacloud/internal/server"
	//alibaba-svc "github.com/aliyun/alibaba-cloud-sdk-go/internal/services"
)

func main() {

	fmt.Printf("REGION_ID : %s\n", os.Getenv("REGION_ID"))
	fmt.Printf("ACCESS_KEY_ID : %s\n", os.Getenv("ACCESS_KEY_ID"))
	fmt.Printf("ACCESS_KEY_SECRET : %s\n", os.Getenv("ACCESS_KEY_SECRET"))

	go func() {
		api.Handler_API()
	}()

	server.Run_Exporter_Server()
}
