package qqenvelop

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// <?xml version="1.0" encoding="UTF-8" ?>
// <xml>
// <appid><![CDATA[1111111111]]></appid>
// <mch_id><![CDATA[1104606907]]></mch_id>
// <openid><![CDATA[JIDJIHUE]]></openid>
// <out_trade_no><![CDATA[10000436560988432048]]></out_trade_no>
// <sign><![CDATA[3B8E2755D32BEE3AF3D45274F03CCB77]]></sign>
// <sign_type><![CDATA[MD5]]></sign_type>
// <state><![CDATA[1]]></state>
// <time_end><![CDATA[2020-01-09 15:21:02]]></time_end>
// <total_fee><![CDATA[1]]></total_fee>
// <transaction_id><![CDATA[29840058602]]></transaction_id>
// </xml>

// NotifyQPayHb notify
type NotifyQPayHb struct {
	AppSecret string
}

// ParseNotify parse
type ParseNotify struct {
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	TransactionID string `xml:"transaction_id"`
	OpenID        string `xml:"openid"`
	TotalFee      int    `xml:"total_fee"`
	ListID        string `xml:"listid"`
	RecvUin       string `xml:"recv_uin"`
	TimeEnd       string `xml:"time_end"`
	State         int    `xml:"state"`
	Sign          string `xml:"sign"`
	SignType      string `xml:"sign_type"`
	Attach        string `xml:"attach"`
}

// Parse parse notify params
func (n *NotifyQPayHb) Parse(b []byte) (*ParseNotify, error) {

	ret := new(ParseNotify)
	if err := xml.Unmarshal(b, ret); err != nil {
		return ret, err
	}
	return ret, nil
}

// Response response
func (n *NotifyQPayHb) Response(code, msg string) string {

	if code == "SUCCESS" {
		return fmt.Sprintf("<xml><return_code>%s</return_code></xml>", "SUCCESS")
	}
	return fmt.Sprintf("<xml><return_code>%s</return_code><return_msg>%s</return_msg></xml>", "FAIL", msg)
}

// CheckSign check sign
func (n *NotifyQPayHb) CheckSign(params *ParseNotify) bool {

	sign := params.Sign

	m := map[string]string{
		"mch_id":         params.MchID,
		"out_trade_no":   params.OutTradeNo,
		"transaction_id": params.TransactionID,
		"openid":         params.OpenID,
		"total_fee":      strconv.Itoa(params.TotalFee),
		"time_end":       params.TimeEnd,
		"state":          strconv.Itoa(params.State),
		"sign_type":      params.SignType,
	}

	if params.AppID != "" {
		m["appid"] = params.AppID
	}
	if params.Attach != "" {
		m["attach"] = params.Attach
	}

	return sign == generateSign(generateQueryStr(n.AppSecret, m))
}
