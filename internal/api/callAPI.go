package api

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	//alibaba-svc "github.com/aliyun/alibaba-cloud-sdk-go/internal/services"

	tool "github.com/siangyeh8818/golang.exporter.alibabacloud/internal/tool"
)

func Handler_API() {
	log.Println("---------------------Handler_API()---------------------")
	for {
		balance := CallApi()
		log.Println("---------------------write balance tp out.csv---------------------")
		tool.WriteWithIoutil("output.csv", balance)
		internal_time, _ := time.ParseDuration(os.Getenv("Get_API_INTERNAL_TIME"))
		//internal_time, _ := strconv.Atoi(os.Getenv("SELEIUM_INTERNAL_TIME"))
		time.Sleep(time.Duration(internal_time))
	}
}

func CallApi() string {
	billclient, err := bssopenapi.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	request := bssopenapi.CreateQueryAccountBalanceRequest()
	request.Scheme = "https"

	response, err := billclient.QueryAccountBalance(request)
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	AvailableAmount := response.Data.AvailableAmount
	//fmt.Print(AvailableAmount)
	AvailableAmount2 := strings.Replace(AvailableAmount, "\n", "", -1)
	AvailableAmount2 = strings.Replace(AvailableAmount2, ",", "", -1)
	//fmt.Print(AvailableAmount2)

	var balanceResulta float64
	//balanceResulta = float64(test)
	balanceResulta, err = strconv.ParseFloat(AvailableAmount2, 64)
	//fmt.Println(test2)
	fmt.Println("------balanceResulta------")
	fmt.Println(balanceResulta)
	return AvailableAmount2
}
