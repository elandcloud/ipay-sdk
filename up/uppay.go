package up

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/pangpanglabs/goutils/httpreq"

	"github.com/fatih/structs"
	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
	"github.com/relax-space/go-kitt/random"
)

func LoopQuery(reqDto ReqQueryDto, custDto ReqCustomerDto, limit, interval int) (int, RespQueryDto, error) {
	if limit <= 0 || limit >= 60 {
		limit = PAY_OVERTIME
	}
	if interval <= 0 || interval >= 5 {
		interval = PAY_LOOPQUERY_INTERVAL
	}
	count := limit / interval
	waitTime := time.Duration(interval) * time.Second //2s
	var status int
	var respQueryDto RespQueryDto
	var err error
	for index := 0; index < count; index++ {
		status, respQueryDto, err = Query(reqDto, custDto)
		if err != nil { //2. request wechat query api failure
			time.Sleep(waitTime)
			continue
		}
		// 1. request wechat query api success
		tradeStatus := respQueryDto.TradeState
		switch {
		case tradeStatus == "SUCCESS": //1.2 pay success
			return status, respQueryDto, nil
		case tradeStatus == "CLOSED" || tradeStatus == "REFUND" || tradeStatus == "REVOKED" || tradeStatus == "REVERSE" || tradeStatus == "NOTPAY" || tradeStatus == "PAYERROR":
			err := fmt.Errorf("pay status is %v", tradeStatus)
			return status, respQueryDto, err
			//1.3 pay failure
		case tradeStatus == "USERPAYING": //1.4 pay unknown
			time.Sleep(waitTime)
			continue
		}
	}
	return status, respQueryDto, err
}

func Pay(reqDto ReqPayDto, custDto ReqCustomerDto) (int, RespPayDto, error) {
	var respDto RespPayDto

	reqDto.Service = "unified.trade.micropay"
	reqDto.Version = "2.0"
	reqDto.Charset = "UTF-8"
	reqDto.SignType = "MD5"
	reqDto.NonceStr = random.Uuid("")

	if len(reqDto.OutTradeNo) == 0 {
		reqDto.OutTradeNo = random.NewUuid(PRE_OUTTRADENO)
	}
	if len(reqDto.MchCreateIp) == 0 {
		reqDto.MchCreateIp = "127.0.0.1"
	}

	reqStruct := structs.New(reqDto)
	reqStruct.TagName = "json"
	reqMap := reqStruct.Map()

	signStr := base.JoinMapObject(reqMap)

	var err error
	reqDto.Sign = sign.MakeMd5Sign(signStr, custDto.Key)
	b, err := xml.MarshalIndent(reqDto, "", " ")
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	xmlData := string(b)
	statusCode, err := httpreq.New(http.MethodPost, OPENAPIURL,
		xmlData, func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.XmlType
			httpReq.RespDataType = httpreq.XmlType
			return nil
		}).
		Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil
}

func Query(reqDto ReqQueryDto, custDto ReqCustomerDto) (int, RespQueryDto, error) {
	var respDto RespQueryDto
	reqDto.Service = "unified.trade.query"
	reqDto.Version = "2.0"
	reqDto.Charset = "UTF-8"
	reqDto.SignType = "MD5"
	reqDto.NonceStr = random.Uuid("")

	reqStruct := structs.New(reqDto)
	reqStruct.TagName = "json"
	reqMap := reqStruct.Map()

	signStr := base.JoinMapObject(reqMap)

	var err error
	reqDto.Sign = sign.MakeMd5Sign(signStr, custDto.Key)
	b, err := xml.MarshalIndent(reqDto, "", " ")
	if err != nil {
		return http.StatusBadRequest, respDto, err
	}
	xmlData := string(b)
	statusCode, err := httpreq.New(http.MethodPost, OPENAPIURL,
		xmlData, func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.XmlType
			httpReq.RespDataType = httpreq.XmlType
			return nil
		}).
		Call(&respDto)
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	return statusCode, respDto, nil
}

func IsQuery(resultCode, errCode string) bool {
	if resultCode != "0" && (errCode != SYSTEMERROR && errCode != BANKERROR && errCode != USERPAYING) {
		return false
	}
	return true
}
