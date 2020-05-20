package api

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	//alibaba-svc "github.com/aliyun/alibaba-cloud-sdk-go/internal/services"
)

func CallApi() float64 {
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
	return balanceResulta
}
