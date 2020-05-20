package api

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	//alibaba-svc "github.com/aliyun/alibaba-cloud-sdk-go/internal/services"
)

func callApi()float64 {
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
	fmt.Println(response.Data.AvailableAmount)
	var balanceResulta float64
	balanceResulta = float64(response.Data.AvailableAmount))
	fmt.Println(balanceResulta)
	return balanceResulta

}