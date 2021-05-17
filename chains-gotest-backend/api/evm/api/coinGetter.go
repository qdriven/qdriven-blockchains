package api

import (
	"github.com/gin-gonic/gin"
	"go-chains/chains/service"
	"go-chains/model/response"
)

type TestCoinRequest struct {
	Address string
}

func GetEthTestCoin(c *gin.Context) {
	request := &TestCoinRequest{}
	_ = c.ShouldBind(&request)
	resp, err := service.GetEthTestCoin(request.Address)
	if err != nil {
		response.FailWithMessage("fail to Get Eth Test Coin", c)
	}
	response.OkWithData(resp, c)
}
