package up

const (
	OPENAPIURL = "https://qra.95516.com/pay/gateway"
)

//https://up.95516.com/open/openapi/doc?index_1=2&index_2=1&chapter_1=275&chapter_2=301
var PayStatusList = []string{
	"SYSTEMERROR",
	"Internal error",
	"BANKERROR",
	"10003",
	"USERPAYING",

	"System error",
	"aop.ACQ.SYSTEM_ERROR",
	"ACQ.SYSTEM_ERROR",
}

const (
	SYSTEMERROR   = "SYSTEMERROR"
	ORDERNOTEXIST = "ORDERNOTEXIST"
	BANKERROR     = "BANKERROR"
	USERPAYING    = "USERPAYING"
)

const (
	MESSAGE_PAYING   = "MESSAGE_PAYING"
	MESSAGE_UNIONPAY = "MESSAGE_UNIONPAY"
	MESSAGE_OVERTIME = "MESSAGE_OVERTIME"
)

const (
	PAY_OVERTIME              = 30 //default 30s
	PAY_LOOPQUERY_INTERVAL    = 2
	REFUND_OVERTIME           = 30 //default 30s
	REFUND_LOOPQUERY_INTERVAL = 2
)

const (
	PRE_OUTTRADENO    = "20"
	PRE_OUTREFUNDNO   = "21"
	PRE_PREOUTTRADENO = "22"
)

func ContainsString(bs []string, b string) bool {
	for _, v := range bs {
		if v == b {
			return true
		}
	}
	return false
}
