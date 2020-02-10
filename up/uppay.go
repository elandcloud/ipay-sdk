package up
import (
	"encoding/xml"
	"net/http"

	"github.com/pangpanglabs/goutils/httpreq"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
	"github.com/relax-space/go-kitt/random"
	"github.com/fatih/structs"
)

func Pay(reqDto ReqPayDto, custDto ReqCustomerDto) (int, RespPayDto, error) {
	var respDto RespPayDto

	BuildCommonParam(&reqDto,"unified.trade.micropay")

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

func BuildCommonParam(reqDto *ReqPayDto, service string)  {
	reqDto.Service=service
	reqDto.Version="2.0"
	reqDto.Charset="UTF-8"
	reqDto.SignType="MD5"
	reqDto.NonceStr=random.Uuid("")
}