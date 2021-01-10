package qqenvelop

import (
	"github.com/birjemin/qqenvelop/utils"
	"log"
)

// qPayHbDetail url
const qPayHbDetail = "https://qpay.qq.com/cgi-bin/mch_query/qpay_hb_mch_list_query.cgi"

// QPayHbDetail detail
type QPayHbDetail struct {
	MchID       string
	AppSecret   string
	HTTPRequest *utils.HTTPClient
}

// ParamsQPayHbDetail params of envelop
type ParamsQPayHbDetail struct {
	ListID    string
	SendType  int    // default 0
	SubMchID  string // detault empty
	MchBillNo string // default 0
}

// RespQPayHbDetail response
type RespQPayHbDetail struct {
	RetCode     string                 `json:"retcode"`
	RetMsg      string                 `json:"retmsg"`
	ListID      string                 `json:"listid"`
	State       string                 `json:"state"`
	TotalNum    string                 `json:"total_num"`
	RecvNum     string                 `json:"recv_num"`
	TotalAmount string                 `json:"total_amount"`
	RecvAmount  string                 `json:"recv_amount"`
	RecvDetails []RespQPayHbRecvDetail `json:"recv_details"`
}

// RespQPayHbRecvDetail detail
type RespQPayHbRecvDetail struct {
	CreateTime string `json:"create_time"`
	Uin        string `json:"uin"`
}

// GetDetail get red envelop detail
// Doc: https://mp.qpay.tenpay.com/buss/wiki/221/2174
func (q *QPayHbDetail) GetDetail(attributes ParamsQPayHbDetail) (*RespQPayHbDetail, error) {
	return q.doGetDetail(qPayHbDetail, attributes)
}

// doGetDetail handle this action
func (q *QPayHbDetail) doGetDetail(URL string, attributes ParamsQPayHbDetail) (*RespQPayHbDetail, error) {

	params := map[string]string{
		"nonce_str": utils.Hex(10),
		"mch_id":    q.MchID,
		"listid":    attributes.ListID,
	}

	// 接收红包者openid或者uin，当re_openid是openid时，qqappid必填
	if attributes.SendType != 0 {
		params["send_type"] = "1"
	}
	if attributes.MchBillNo != "" {
		params["mch_billno"] = attributes.MchBillNo
	}
	if attributes.SubMchID != "" {
		params["sub_mch_id"] = attributes.SubMchID
	}

	params["sign"] = generateSign(generateQueryStr(q.AppSecret, params))

	var resp = new(RespQPayHbDetail)

	if err := q.HTTPRequest.HTTPGet(URL, params); err != nil {
		log.Println("[detail]do, post failed", err)
		return resp, err
	}

	if err := q.HTTPRequest.GetResponseJSON(resp); err != nil {
		log.Println("[detail]do, response json failed", err)
		return resp, err
	}
	return resp, nil
}
