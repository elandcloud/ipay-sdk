package tests

import (
	"fmt"
	"lemon-ipay-sdk/up"
	"net/http"
	"os"
	"testing"

	"github.com/pangpanglabs/goutils/test"
	"github.com/relax-space/go-kitt/random"
)

var mchId = os.Getenv("UP_MCHID")
var upKey = os.Getenv("UP_KEY")

// go test -run TestUpPay
func TestUpPay(t *testing.T) {
	outTradeNo := random.NewUuid(up.PRE_OUTTRADENO)
	reqDto := up.ReqPayDto{
		MchId:      mchId,
		AuthCode:   "134540965666507931",
		Body:       "xinmiao test up",
		TotalFee:   1,
		OutTradeNo: outTradeNo,
	}
	customDto := up.ReqCustomerDto{
		Key: upKey,
	}

	_, resp, err := up.Pay(reqDto, customDto)
	test.Ok(t, err)
	if up.IsQuery(resp.ResultCode, resp.ErrCode) == false {
		fmt.Println(resp.ErrMsg)
		return
	}

	reqQueryDto := up.ReqQueryDto{
		MchId:      mchId,
		OutTradeNo: outTradeNo,
	}

	status, respQuery, err := up.LoopQuery(reqQueryDto, customDto, 60, 0)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, status)
	test.Equals(t, outTradeNo, respQuery.OutTradeNo)
	test.Equals(t, 1, respQuery.TotalFee)
	test.Equals(t, 1, respQuery.CashFee)
}

//go test -run TestUpQuery
func TestUpQuery(t *testing.T) {
	expDto := up.RespQueryDto{
		MchId:            "QRA290454410UV5",
		TradeState:       "SUCCESS",
		Appid:            "wxc6634b0d53a240d5",
		TransactionId:    "9551600000202002114255187210",
		OutTransactionId: "4200000475202002117266214187",

		OutTradeNo: "202002111767868850625397773",
		TotalFee:   1,
		CashFee:    1,
		TimeEnd:    "20200211233307",
	}
	reqDto := up.ReqQueryDto{
		MchId:      mchId,
		OutTradeNo: "202002111767868850625397773",
	}
	customDto := up.ReqCustomerDto{
		Key: upKey,
	}
	status, resp, err := up.Query(reqDto, customDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, status)
	test.Equals(t, expDto.MchId, resp.MchId)
	test.Equals(t, expDto.TradeState, resp.TradeState)
	test.Equals(t, expDto.Appid, resp.Appid)
	test.Equals(t, expDto.TransactionId, resp.TransactionId)
	test.Equals(t, expDto.OutTransactionId, resp.OutTransactionId)

	test.Equals(t, expDto.OutTradeNo, resp.OutTradeNo)
	test.Equals(t, expDto.TotalFee, resp.TotalFee)
	test.Equals(t, expDto.CashFee, resp.CashFee)
	test.Equals(t, expDto.TimeEnd, resp.TimeEnd)
}
