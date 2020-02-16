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

	// for good detail
	CostPrice int    `xml:"cost_price,omitempty" json:"cost_price,omitempty"`
	ReceiptId string `xml:"receipt_id,omitempty" json:"receipt_id,omitempty"`
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

	TradeState     string `xml:"trade_state,omitempty" json:"trade_state,omitempty"`
	TradeStateDesc string `xml:"trade_state_desc,omitempty" json:"trade_state_desc,omitempty"`
	TradeType      string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	Appid          string `xml:"appid,omitempty" json:"appid,omitempty"`
	SubAppid       string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty"`
	Openid         string `xml:"openid,omitempty" json:"openid,omitempty"`
	SubOpenid      string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"`
	IsSubscribe    string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	SubIsSubscribe string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty"`

	TransactionId    string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTransactionId string `xml:"out_transaction_id,omitempty" json:"out_transaction_id,omitempty"`
	OutTradeNo       string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	TotalFee         int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFee          int    `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`

	PromotionDetail  string  `xml:"promotion_detail,omitempty" json:"promotion_detail,omitempty"`
	UnionpayDiscount string  `xml:"unionpay_discount,omitempty" json:"unionpay_discount,omitempty"`
	InvoiceAmount    float32 `xml:"invoice_amount,omitempty" json:"invoice_amount,omitempty"`
	BuyerLogonId     string  `xml:"buyer_logon_id,omitempty" json:"buyer_logon_id,omitempty"`
	BuyerPayAmount   float32 `xml:"buyer_pay_amount,omitempty" json:"buyer_pay_amount,omitempty"`

	BuyerUserId         string  `xml:"buyer_user_id,omitempty" json:"buyer_user_id,omitempty"`
	PointAmount         float32 `xml:"point_amount,omitempty" json:"point_amount,omitempty"`
	ReceiptAmount       float32 `xml:"receipt_amount,omitempty" json:"receipt_amount,omitempty"`
	FundBillList        string  `xml:"fund_bill_list,omitempty" json:"fund_bill_list,omitempty"`
	DiscountGoodsDetail string  `xml:"discount_goods_detail,omitempty" json:"discount_goods_detail,omitempty"`

	CouponFee int    `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty"`
	Mdiscount int    `xml:"mdiscount,omitempty" json:"mdiscount,omitempty"`
	FeeType   string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	Attach    string `xml:"attach,omitempty" json:"attach,omitempty"`
	BankType  string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`

	BankBillno string `xml:"bank_billno,omitempty" json:"bank_billno,omitempty"`
	TimeEnd    string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	SettleKey  string `xml:"settle_key,omitempty" json:"settle_key,omitempty"`
}

// 3. refund

type ReqRefundDto struct {
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
	OutRefundNo   string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	TotalFee      int    `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	RefundFee     int    `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`

	OpUserId      string `xml:"op_user_id,omitempty" json:"op_user_id,omitempty"`
	RefundChannel string `xml:"refund_channel,omitempty" json:"refund_channel,omitempty"`
}

type RespRefundDto struct {
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

	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	OutRefundNo   string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	RefundId      string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
	RefundChannel string `xml:"refund_channel,omitempty" json:"refund_channel,omitempty"`

	RefundFee       int    `xml:"refund_fee,omitempty" json:"refund_fee,omitempty"`
	CouponRefundFee int    `xml:"coupon_refund_fee,omitempty" json:"coupon_refund_fee,omitempty"`
	TradeType       string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
}

// 4. refund query

type ReqRefundQueryDto struct {
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
	OutRefundNo   string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`
	RefundId      string `xml:"refund_id,omitempty" json:"refund_id,omitempty"`
}

type RespRefundQueryDto struct {
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

	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	RefundCount   int    `xml:"refund_count,omitempty" json:"refund_count,omitempty"`
	TradeType     string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`

	//Because Alipay can only query a single refund, in order to unified format, all refunds only return one data
	OutRefundNo     string `xml:"out_refund_no_0,omitempty" json:"out_refund_no_0,omitempty"`
	RefundId        string `xml:"refund_id_0,omitempty" json:"refund_id_0,omitempty"`
	RefundChannel   string `xml:"refund_channel_0,omitempty" json:"refund_channel_0,omitempty"`
	RefundFee       int    `xml:"refund_fee_0,omitempty" json:"refund_fee_0,omitempty"`
	CouponRefundFee int    `xml:"coupon_refund_fee_0,omitempty" json:"coupon_refund_fee_0,omitempty"`

	Mdiscount    int    `xml:"mdiscount_0,omitempty" json:"mdiscount_0,omitempty"`
	RefundTime   string `xml:"refund_time_0,omitempty" json:"refund_time_0,omitempty"`
	RefundStatus string `xml:"refund_status_0,omitempty" json:"refund_status_0,omitempty"`
	SettleKey    string `xml:"settle_key_0,omitempty" json:"settle_key_0,omitempty"`
}

// custom

type ReqCustomerDto struct {
	Url string `json:"url,omitempty"`
	Key string `json:"key,omitempty"`
}

// sub dto
type GoodsDetailReqDto struct {
	GoodsId    string  `xml:"goods_id,omitempty" json:"goods_id,omitempty"`
	PayGoodsId string  `xml:"pay_goods_id,omitempty" json:"pay_goods_id,omitempty"`
	GoodsName  string  `xml:"goods_name,omitempty" json:"goods_name,omitempty"`
	Quantity   int     `xml:"quantity,omitempty" json:"quantity,omitempty"`
	Price      float32 `xml:"price,omitempty" json:"price,omitempty"`

	GoodsCategory string `xml:"goods_category,omitempty" json:"goods_category,omitempty"`
	Body          string `xml:"body,omitempty" json:"body,omitempty"`
	ShowUrl       string `xml:"show_url,omitempty" json:"show_url,omitempty"`
	Version       string `xml:"version,omitempty" json:"version,omitempty"`
}

type DiscountGoodsDetailRespDto struct {
	GoodsId        string `xml:"goods_id,omitempty" json:"goods_id,omitempty"`
	GoodsName      string `xml:"goods_name,omitempty" json:"goods_name,omitempty"`
	DiscountAmount string `xml:"discount_amount,omitempty" json:"discount_amount,omitempty"`
	VoucherId      string `xml:"voucher_id,omitempty" json:"voucher_id,omitempty"`
}

type PromotionDetailRespDto struct {
	PromotionId string `xml:"promotion_id,omitempty" json:"promotion_id,omitempty"`
	Name        string `xml:"name,omitempty" json:"name,omitempty"`
	Scope       string `xml:"scope,omitempty" json:"scope,omitempty"`
	Type        string `xml:"type,omitempty" json:"type,omitempty"`
	Amount      int    `xml:"amount,omitempty" json:"amount,omitempty"`

	ActivityId         string               `xml:"activity_id,omitempty" json:"activity_id,omitempty"`
	WxpayContribute    int                  `xml:"wxpay_contribute,omitempty" json:"wxpay_contribute,omitempty"`
	MerchantContribute int                  `xml:"merchant_contribute,omitempty" json:"merchant_contribute,omitempty"`
	OtherContribute    int                  `xml:"other_contribute,omitempty" json:"other_contribute,omitempty"`
	GoodsDetail        []GoodsDetailRespDto `xml:"goods_detail,omitempty" json:"goods_detail,omitempty"`
}

type GoodsDetailRespDto struct {
	GoodsId         string `xml:"goods_id,omitempty" json:"goods_id,omitempty"`
	Quantity        int    `xml:"quantity,omitempty" json:"quantity,omitempty"`
	Price           int    `xml:"price,omitempty" json:"price,omitempty"`
	discount_amount int    `xml:"discount_amount,omitempty" json:"discount_amount,omitempty"`
}
