package qqenvelop

import (
	"encoding/xml"
	"fmt"
)

// <xml>
//   <mch_id><![CDATA[1104606907]]></mch_id>
//   <mch_billno><![CDATA[29840058602]]></mch_billno>
//   <listid><![CDATA[10000436560988432048]]></listid>
//   <recv_uin><![CDATA[2344546]]></recv_uin>
//   <total_fee><![CDATA[10]]></total_fee>
//   <time_end><![CDATA[20161025094946]]></time_end>
//   <state><![CDATA[1]]></state>
//   <sign><![CDATA[DE4335434F33C065C449E261DCE08BCF]]></sign>
//   <sign_type><![CDATA[MD5]]></sign_type>
// </xml>

// NotifyQPayHb notify
type NotifyQPayHb struct {
}

// ParseNotify parse
type ParseNotify struct {
	AppID     string `xml:"appid"`
	MchID     string `xml:"mch_id"`
	MchBillNo string `xml:"mch_billno"`
	ListID    string `xml:"listid"`
	RecvUin   string `xml:"recv_uin"`
	OpenID    string `xml:"openid"`
	TotalFee  int    `xml:"total_fee"`
	TimeEnd   string `xml:"time_end"`
	State     int    `xml:"state"`
	Attach    string `xml:"attach"`
	Sign      string `xml:"sign"`
	SignType  string `xml:"sign_type"`
}

// Parse parse notify params
func (r NotifyQPayHb) Parse(b []byte) (*ParseNotify, error) {

	ret := new(ParseNotify)
	if err := xml.Unmarshal(b, ret); err != nil {
		return ret, err
	}
	return ret, nil
}

// Response response
func (r NotifyQPayHb) Response(code, msg string) string {

	if code == "SUCCESS" {
		return fmt.Sprintf("<xml><return_code>%s</return_code></xml>", "SUCCESS")
	}
	return fmt.Sprintf("<xml><return_code>%s</return_code><return_msg>%s</return_msg></xml>", "FAIL", msg)
}
