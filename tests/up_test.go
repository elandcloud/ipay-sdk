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

//{XMLName:{Space: Local:xml} Version:2.0 Charset:UTF-8 SignType:MD5 SignAgentno: Groupno: Status:0 Message: Code: ResultCode:0 MchId:QRA290454410UV5 DeviceInfo: NonceStr:3536582798494089883 ErrCode: ErrMsg: Sign:FAE876E8C337CA1F199325C781B00456 TransactionId:9551600000202002114255187210 OutTradeNo:202002111767868850625397773 OutRefundNo:212002132220972889823986055 RefundId:9551600000202002133245848244 RefundChannel:ORIGINAL RefundFee:1 CouponRefundFee:0 TradeType:pay.weixin.micropay}

//go test -run TestUpRefund
func TestUpRefund(t *testing.T) {
	expDto := up.RespRefundDto{
		Status:            "0",
	}
	outRefundNo := random.NewUuid(up.PRE_OUTREFUNDNO)
	reqDto := up.ReqRefundDto{
		MchId:      mchId,
		OutTradeNo: "202002111767868850625397773",
		TotalFee:   1,
		RefundFee:1,
		OutRefundNo:outRefundNo,
	}
	customDto := up.ReqCustomerDto{
		Key: upKey,
	}
	status, resp, err := up.Refund(reqDto, customDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, status)
	test.Equals(t, expDto.Status, resp.Status)
}