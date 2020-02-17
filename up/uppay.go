package up

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pangpanglabs/goutils/httpreq"

	"github.com/fatih/structs"
	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
	"github.com/relax-space/go-kitt/random"
)

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
	resp, err := httpreq.New(http.MethodPost, OPENAPIURL,
		xmlData, func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.XmlType
			return nil
		}).
		RawCall()
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusOK, respDto, err
	}
	if CheckSign(b, custDto.Key) == false {
		return http.StatusOK, respDto, errors.New("response sign error")
	}

	if err := xml.Unmarshal(b, &respDto); err != nil {
		return http.StatusOK, respDto, errors.New("response parse error")
	}

	return resp.StatusCode, respDto, nil
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
	resp, err := httpreq.New(http.MethodPost, OPENAPIURL,
		xmlData, func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.XmlType
			return nil
		}).
		RawCall()
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusOK, respDto, err
	}
	if CheckSign(b, custDto.Key) == false {
		return http.StatusOK, respDto, errors.New("response sign error")
	}

	if err := xml.Unmarshal(b, &respDto); err != nil {
		return http.StatusOK, respDto, errors.New("response parse error")
	}

	return resp.StatusCode, respDto, nil
}

func Refund(reqDto ReqRefundDto, custDto ReqCustomerDto) (int, RespRefundDto, error) {
	var respDto RespRefundDto

	reqDto.Service = "unified.trade.refund"
	reqDto.Version = "2.0"
	reqDto.Charset = "UTF-8"
	reqDto.SignType = "MD5"
	reqDto.NonceStr = random.Uuid("")

	if len(reqDto.OutRefundNo) == 0 {
		reqDto.OutRefundNo = random.NewUuid(PRE_OUTREFUNDNO)
	}

	if len(reqDto.OpUserId) == 0 {
		reqDto.OpUserId = reqDto.MchId
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
	resp, err := httpreq.New(http.MethodPost, OPENAPIURL,
		xmlData, func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.XmlType
			return nil
		}).
		RawCall()
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusOK, respDto, err
	}
	if CheckSign(b, custDto.Key) == false {
		return http.StatusOK, respDto, errors.New("response sign error")
	}

	if err := xml.Unmarshal(b, &respDto); err != nil {
		return http.StatusOK, respDto, errors.New("response parse error")
	}

	return resp.StatusCode, respDto, nil
}

func RefundQuery(reqDto ReqRefundQueryDto, custDto ReqCustomerDto) (int, RespRefundQueryDto, error) {
	var respDto RespRefundQueryDto

	reqDto.Service = "unified.trade.refundquery"
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
	resp, err := httpreq.New(http.MethodPost, OPENAPIURL,
		xmlData, func(httpReq *httpreq.HttpReq) error {
			httpReq.ReqDataType = httpreq.XmlType
			return nil
		}).
		RawCall()
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return http.StatusInternalServerError, respDto, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusOK, respDto, err
	}
	if CheckSign(b, custDto.Key) == false {
		return http.StatusOK, respDto, errors.New("response sign error")
	}

	if err := xml.Unmarshal(b, &respDto); err != nil {
		return http.StatusOK, respDto, errors.New("response parse error")
	}

	return resp.StatusCode, respDto, nil
}

func IsQuery(resultCode, errCode string) bool {

	if resultCode == "0" {
		return true
	}
	if ContainsString(PayStatusList, errCode) {
		return true
	}
	return false
}

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
		if err != nil { //2. request  query api failure
			time.Sleep(waitTime)
			continue
		}
		// 1. request  query api success
		tradeStatus := respQueryDto.TradeState
		switch {
		case tradeStatus == "SUCCESS": //1.2 pay success
			return status, respQueryDto, nil
		case tradeStatus == "CLOSED" || tradeStatus == "REFUND" || tradeStatus == "REVOKED" || tradeStatus == "REVERSE" || tradeStatus == "NOTPAY" || tradeStatus == "PAYERROR": //1.3 pay failure
			err := fmt.Errorf("pay status is %v", tradeStatus)
			return status, respQueryDto, err
		case tradeStatus == "USERPAYING": //1.4 pay unknown
			time.Sleep(waitTime)
			continue
		}
	}
	return status, respQueryDto, err
}

func LoopRefundQuery(reqDto ReqRefundQueryDto, custDto ReqCustomerDto, limit, interval int) (int, RespRefundQueryDto, error) {
	if limit <= 0 || limit >= 60 {
		limit = REFUND_OVERTIME
	}
	if interval <= 0 || interval >= 5 {
		interval = REFUND_LOOPQUERY_INTERVAL
	}
	count := limit / interval
	waitTime := time.Duration(interval) * time.Second //2s
	var status int
	var respRefundQueryDto RespRefundQueryDto
	var err error
	for index := 0; index < count; index++ {
		status, respRefundQueryDto, err = RefundQuery(reqDto, custDto)
		if err != nil { //2. request  query api failure
			time.Sleep(waitTime)
			continue
		}
		// 1. request query api success
		tradeStatus := respRefundQueryDto.RefundStatus
		switch {
		case tradeStatus == "SUCCESS": //1.2 refund success
			return status, respRefundQueryDto, nil
		case tradeStatus == "FAIL" || tradeStatus == "CHANGE": //1.3 refund failure
			err := fmt.Errorf("refund status is %v", tradeStatus)
			return status, respRefundQueryDto, err
		case tradeStatus == "PROCESSING" || tradeStatus == "NOTSURE": //1.4 refund unknown
			time.Sleep(waitTime)
			continue
		}
	}
	return status, respRefundQueryDto, err
}

func ParseDiscountGoodsDetail(str string) (DiscountGoodsDetailRespDto, error) {
	var dto DiscountGoodsDetailRespDto
	err := json.Unmarshal([]byte(str), &dto)
	if err != nil {
		return dto, err
	}
	return dto, nil
}

func ParsePromotionDetail(str string) (PromotionDetailRespDto, error) {
	var dto PromotionDetailRespDto
	err := json.Unmarshal([]byte(str), &dto)
	if err != nil {
		return dto, err
	}
	return dto, nil
}

func XmlToMap(xmlStr string) (map[string]interface{}, error) {
	data, err := base.XmlToMap(xmlStr)
	if err != nil {
		return map[string]interface{}{}, err
	}
	v, has := data["xml"]
	if has == false {
		return map[string]interface{}{}, errors.New("parse resp failure")
	}
	return v.(map[string]interface{}), nil
}

func CheckSign(b []byte, key string) bool {
	xmlStr := string(b)
	respMap, err := XmlToMap(xmlStr)
	if err != nil {
		return false
	}
	signObj, ok := respMap["sign"]
	if ok == false {
		return false
	}
	actSign := signObj.(string)
	delete(respMap, "sign")
	signStr := base.JoinMapObject(respMap)
	expSign := sign.MakeMd5Sign(signStr, key)
	if actSign != expSign {
		return false
	}
	return true
}
