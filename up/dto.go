package up

import (
	"encoding/xml"
)

// 1.pay

type ReqPayDto struct {
	XMLName xml.Name `xml:"xml" json:"xml,omitempty"`

	//========== common

	Service  string `xml:"service,omitempty" json:"service,omitempty"`
	Version  string `xml:"version,omitempty" json:"version,omitempty"`
	Charset  string `xml:"charset,omitempty" json:"charset,omitempty"`
	SignType string `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	MchId    string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`

	NonceStr    string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign        string `xml:"sign,omitempty" json:"sign,omitempty"`
	SignAgentno string `xml:"sign_agentno,omitempty" json:"sign_agentno,omitempty"`
	Groupno     string `xml:"groupno,omitempty" json:"groupno,omitempty"`

	//========== pay req

	OutTradeNo  string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	DeviceInfo  string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	Body        string `xml:"body,omitempty" json:"body,omitempty"`
	GoodsDetail string `xml:"goods_detail,omitempty" json:"goods_detail,omitempty"`
	SubAppId    string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	Attach      string `xml:"attach,omitempty" json:"attach,omitempty"`

	NeedReceipt bool   `xml:"need_receipt,omitempty" json:"need_receipt,omitempty"`
	TotalFee    int64  `xml:"total_fee,omitempty" json:"total_fee,omitempty"` //int64
	MchCreateIp string `xml:"mch_create_ip,omitempty" json:"mch_create_ip,omitempty"`
	AuthCode    string `xml:"auth_code,omitempty" json:"auth_code,omitempty"` //10,11,12,13,14,15
	TimeStart   string `xml:"time_start,omitempty" json:"time_start,omitempty"`

	TimeExpire string `xml:"time_expire,omitempty" json:"time_expire,omitempty"`
	OpUserId   string `xml:"op_user_id,omitempty" json:"op_user_id,omitempty"`
	OpShopId   string `xml:"op_shop_id,omitempty" json:"op_shop_id,omitempty"`
	OpDeviceId string `xml:"op_device_id,omitempty" json:"op_device_id,omitempty"`
	GoodsTag   string `xml:"goods_tag,omitempty" json:"goods_tag,omitempty"`
}

//response

type RespPayDto struct {
	XMLName xml.Name `xml:"xml" json:"xml,omitempty"`

	//========== common

	Version     string `xml:"version,omitempty" json:"version,omitempty"`
	Charset     string `xml:"charset,omitempty" json:"charset,omitempty"`
	SignType    string `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	SignAgentno string `xml:"sign_agentno,omitempty" json:"sign_agentno,omitempty"`
	Groupno     string `xml:"groupno,omitempty" json:"groupno,omitempty"`

	Status    string `xml:"status,omitempty" json:"status,omitempty"`
	Message   string `xml:"message,omitempty" json:"message,omitempty"`
	Code      string `xml:"code,omitempty" json:"code,omitempty"`
	NeedQuery string `xml:"need_query,omitempty" json:"need_query,omitempty"`

	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`

	ErrMsg string `xml:"err_msg,omitempty" json:"err_msg,omitempty"`
	Sign   string `xml:"sign,omitempty" json:"sign,omitempty"`

	//========== pay resp

	Appid       string `xml:"appid,omitempty" json:"appid,omitempty"`
	Openid      string `xml:"openid,omitempty" json:"openid,omitempty"`
	SubOpenid   string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`
	TradeType   string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	IsSubscribe string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`

	PayResult        int    `xml:"pay_result,omitempty" json:"pay_result,omitempty"`
	PayInfo          string `xml:"pay_info,omitempty" json:"pay_info,omitempty"`
	TransactionId    string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTransactionId string `xml:"out_transaction_id,omitempty" json:"out_transaction_id,omitempty"`
	SubIsSubscribe   string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty"`

	SubAppid   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	OutTradeNo string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	TotalFee   int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFee    int    `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	Mdiscount  int    `xml:"mdiscount,omitempty" json:"mdiscount,omitempty"`

	InvoiceAmount   string          `xml:"invoice_amount,omitempty" json:"invoice_amount,omitempty"`
	CouponFee       int             `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	PromotionDetail string          `xml:"promotion_detail,omitempty" json:"promotion_detail,omitempty"`
	FundBillList    FundBillListDto `xml:"fund_bill_list,omitempty" json:"fund_bill_list,omitempty"`
	BuyerLogonId    string          `xml:"buyer_logon_id,omitempty" json:"buyer_logon_id,omitempty"`

	BuyerPayAmount      float32 `xml:"buyer_pay_amount,omitempty" json:"buyer_pay_amount,omitempty"`
	BuyerUserId         string  `xml:"buyer_user_id,omitempty" json:"buyer_user_id,omitempty"`
	PointAmount         float32 `xml:"point_amount,omitempty" json:"point_amount,omitempty"`
	ReceiptAmount       float32 `xml:"receipt_amount,omitempty" json:"receipt_amount,omitempty"`
	DiscountGoodsDetail string  `xml:"discount_goods_detail,omitempty" json:"discount_goods_detail,omitempty"`

	FeeType    string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	Attach     string `xml:"attach,omitempty" json:"attach,omitempty"`
	BankType   string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	Uuid       string `xml:"uuid,omitempty" json:"uuid,omitempty"`
	BankBillno string `xml:"bank_billno,omitempty" json:"bank_billno,omitempty"`

	TimeEnd string `xml:"time_end,omitempty" json:"time_end,omitempty"`
}

type FundBillListDto struct {
	fund_channel string  `xml:"fund_channel,omitempty" json:"fund_channel,omitempty"`
	amount       float32 `xml:"amount,omitempty" json:"amount,omitempty"`
}

// 2.query

type ReqQueryDto struct {
	XMLName xml.Name `xml:"xml" json:"xml,omitempty"`

	//========== common

	Service  string `xml:"service,omitempty" json:"service,omitempty"`
	Version  string `xml:"version,omitempty" json:"version,omitempty"`
	Charset  string `xml:"charset,omitempty" json:"charset,omitempty"`
	SignType string `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	MchId    string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`

	NonceStr    string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign        string `xml:"sign,omitempty" json:"sign,omitempty"`
	SignAgentno string `xml:"sign_agentno,omitempty" json:"sign_agentno,omitempty"`
	Groupno     string `xml:"groupno,omitempty" json:"groupno,omitempty"`

	//========== pay req

	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
}

type RespQueryDto struct {
	XMLName xml.Name `xml:"xml" json:"xml,omitempty"`

	//========== common

	Version     string `xml:"version,omitempty" json:"version,omitempty"`
	Charset     string `xml:"charset,omitempty" json:"charset,omitempty"`
	SignType    string `xml:"sign_type,omitempty" json:"sign_type,omitempty"`
	SignAgentno string `xml:"sign_agentno,omitempty" json:"sign_agentno,omitempty"`
	Groupno     string `xml:"groupno,omitempty" json:"groupno,omitempty"`

	Status     string `xml:"status,omitempty" json:"status,omitempty"`
	Message    string `xml:"message,omitempty" json:"message,omitempty"`
	Code       string `xml:"code,omitempty" json:"code,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`

	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrMsg     string `xml:"err_msg,omitempty" json:"err_msg,omitempty"`

	Sign string `xml:"sign,omitempty" json:"sign,omitempty"`

	//========== pay resp

	TradeState       string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	TradeStateDesc       string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty"`
	TradeType       string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	Appid       string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid       string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	Openid      string `xml:"openid,omitempty" json:"openid,omitempty"`
	SubOpenid   string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`
	IsSubscribe string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	SubIsSubscribe string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty"`

	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTransactionId string `xml:"out_transaction_id,omitempty" json:"out_transaction_id,omitempty"`
	OutTradeNo string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	TotalFee int `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFee int `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`

	PromotionDetail string `xml:"promotion_detail,omitempty" json:"promotion_detail,omitempty"`
	UnionpayDiscount string `xml:"unionpay_discount,omitempty" json:"unionpay_discount,omitempty"`
	InvoiceAmount  float32 `xml:"invoice_amount,omitempty" json:"invoice_amount,omitempty"`
	BuyerLogonId string `xml:"buyer_logon_id,omitempty" json:"buyer_logon_id,omitempty"`
	BuyerPayAmount string `xml:"buyer_pay_amount,omitempty" json:"buyer_pay_amount,omitempty"`
	
	BuyerUserId string `xml:"buyer_user_id,omitempty" json:"buyer_user_id,omitempty"`
	PointAmount float32 `xml:"point_amount,omitempty" json:"point_amount,omitempty"`
	ReceiptAmount float32 `xml:"receipt_amount,omitempty" json:"receipt_amount,omitempty"`
	FundBillList string `xml:"fund_bill_list,omitempty" json:"fund_bill_list,omitempty"`
	DiscountGoodsDetail string `xml:"discount_goods_detail,omitempty" json:"discount_goods_detail,omitempty"`

	CouponFee      int `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	Mdiscount      int `xml:"mdiscount,omitempty" json:"mdiscount,omitempty"`
	FeeType      string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	Attach      string `xml:"attach,omitempty" json:"attach,omitempty"`
	BankType      string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`

	BankBillno      string `xml:"bank_billno,omitempty" json:"bank_billno,omitempty"`
	TimeEnd      string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	SettleKey      string `xml:"settle_key,omitempty" json:"settle_key,omitempty"`


}

// custom

type ReqCustomerDto struct {
	Url string `json:"url,omitempty"`
	Key string `json:"key,omitempty"`
}
