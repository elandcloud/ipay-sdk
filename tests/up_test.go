package tests

import (
	"fmt"
	"lemon-ipay-sdk/up"
	"net/http"
	"os"
	"testing"
	"encoding/json"

	"github.com/pangpanglabs/goutils/test"
	"github.com/relax-space/go-kitt/random"
)

var mchId = os.Getenv("UP_MCHID")
var upKey = os.Getenv("UP_KEY")

/**
{
    "xml": {
        "Space": "", 
        "Local": "xml"
    }, 
    "version": "2.0", 
    "charset": "UTF-8", 
    "sign_type": "MD5", 
    "status": "0", 
    "result_code": "0", 
    "mch_id": "QRA290454410UV5", 
    "nonce_str": "9928960196914791592", 
    "sign": "2AC6C72CEA46F54A2BDC5C369AD9D464", 
    "trade_state": "SUCCESS", 
    "trade_type": "pay.alipay.micropay", 
    "appid": "1266000004184930", 
    "openid": "2088702305824122", 
    "transaction_id": "9551600000202002163247543073", 
    "out_transaction_id": "2020021622001424121420648297", 
    "out_trade_no": "2020021611965285238855190655", 
    "total_fee": 100, 
    "invoice_amount": 1, 
    "buyer_logon_id": "xia***@163.com", 
    "buyer_pay_amount": 1, 
    "buyer_user_id": "2088702305824122", 
    "receipt_amount": 1, 
    "fund_bill_list": "[{\"amount\":\"1.00\",\"fundChannel\":\"ALIPAYACCOUNT\"}]", 
    "fee_type": "CNY", 
    "bank_type": "ALIPAYACCOUNT", 
    "time_end": "20200216203722"
}

**/
// go test -run TestUpPay
func TestUpPay(t *testing.T) {
	goodsDetail :=up.GoodsDetailReqDto{
		GoodsId:"8054489",
		GoodsName:"yp8054489",
		Quantity:1,
		Price:100,
	}
	b,_:=json.Marshal(goodsDetail)
	goodsDetailStr:=string(b)
	
	outTradeNo := random.NewUuid(up.PRE_OUTTRADENO)
	reqDto := up.ReqPayDto{
		MchId:      mchId,
		AuthCode:   "28787383905994393",
		Body:       "xinmiao test up",
		TotalFee:   1,
		OutTradeNo: outTradeNo,
		GoodsDetail:goodsDetailStr,
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
	// fmt.Printf("%+v",respQuery)
	test.Equals(t, outTradeNo, respQuery.OutTradeNo)
	test.Equals(t, 1, respQuery.TotalFee)
	if len(respQuery.PromotionDetail) !=0{
		_,err=up.ParseDiscountGoodsDetail(respQuery.PromotionDetail)
		test.Ok(t,err)
	}
	if len(respQuery.DiscountGoodsDetail) !=0{
		_,err=up.ParsePromotionDetail(respQuery.DiscountGoodsDetail)
		test.Ok(t,err)
	}
}

//go test -run TestUpQuery
func TestUpQuery(t *testing.T) {
	expDto := up.RespQueryDto{
		MchId:            "QRA290454410UV5",
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
	outTradeNo :="2020021611965285238855190655"
	outRefundNo := random.NewUuid(up.PRE_OUTREFUNDNO)
	refundFee:=1
	totalFee:=100
	expDto := up.RespRefundQueryDto{
		Status:            "0",
		OutTradeNo: outTradeNo,
		OutRefundNo: outRefundNo,
		RefundFee:refundFee,
	}
	reqDto := up.ReqRefundDto{
		MchId:      mchId,
		OutTradeNo: outTradeNo,
		TotalFee:   totalFee,
		RefundFee:refundFee,
		OutRefundNo:outRefundNo,
	}
	customDto := up.ReqCustomerDto{
		Key: upKey,
	}
	_, _, err := up.Refund(reqDto, customDto)
	test.Ok(t, err)

	reqRefundQueryDto := up.ReqRefundQueryDto{
		MchId:      mchId,
		OutTradeNo: outTradeNo,
		OutRefundNo: outRefundNo,
	}

	status, respRefundQuery, err := up.LoopRefundQuery(reqRefundQueryDto, customDto, 60, 0)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, status)
	// fmt.Printf("%+v",respRefundQuery)
	test.Equals(t, expDto.Status, respRefundQuery.Status)
	test.Equals(t, expDto.OutTradeNo, respRefundQuery.OutTradeNo)
	test.Equals(t, expDto.OutRefundNo, respRefundQuery.OutRefundNo)
	test.Equals(t, expDto.RefundFee, respRefundQuery.RefundFee)

}


//go test -run TestUpRefundQuery
func TestUpRefundQuery(t *testing.T) {
	expDto := up.RespRefundQueryDto{
		MchId:            "QRA290454410UV5",
		TransactionId:    "9551600000202002163247444596",
		OutTradeNo: "202002168827092713337078464",
		TradeType:"pay.alipay.micropay",
		OutRefundNo:"212002164288363517991342006",

		RefundId:"9551600000202002163247447424",
		RefundFee:1,
		RefundTime:"20200216173243",
		RefundStatus:"SUCCESS",
	}
	reqDto := up.ReqRefundQueryDto{
		MchId:      mchId,
		OutTradeNo: "202002168827092713337078464",
		OutRefundNo:"212002164288363517991342006",
	}
	customDto := up.ReqCustomerDto{
		Key: upKey,
	}
	status, resp, err := up.RefundQuery(reqDto, customDto)
	test.Ok(t, err)
	test.Equals(t, http.StatusOK, status)
	test.Equals(t, expDto.MchId, resp.MchId)
	test.Equals(t, expDto.TransactionId, resp.TransactionId)
	test.Equals(t, expDto.OutTradeNo, resp.OutTradeNo)
	test.Equals(t, expDto.TradeType, resp.TradeType)
	test.Equals(t, expDto.OutRefundNo, resp.OutRefundNo)

	test.Equals(t, expDto.RefundId, resp.RefundId)
	test.Equals(t, expDto.RefundFee, resp.RefundFee)
	test.Equals(t, expDto.RefundTime, resp.RefundTime)
	test.Equals(t, expDto.RefundStatus, resp.RefundStatus)
}

