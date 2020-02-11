package tests
import (
	"fmt"
	"lemon-ipay-sdk/up"
	"net/http"
	"os"
	"testing"

	"github.com/pangpanglabs/goutils/test"
)
var mchId  =os.Getenv("UP_MCHID")
var upKey  =os.Getenv("UP_KEY")

func TestUpPay(t *testing.T) {
	reqDto := up.ReqPayDto{
		MchId: mchId,
		AuthCode:    "134981593479354294",
		Body:     "xinmiao test up",
		TotalFee: 1,
	}
	customDto := up.ReqCustomerDto{
		Key: upKey,
	}
	status, resp, err := up.Pay(reqDto, customDto)
	test.Ok(t, err)
	test.Equals(t,http.StatusOK,status)
	fmt.Printf("%+v",resp)
}