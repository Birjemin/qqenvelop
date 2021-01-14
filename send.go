package qqenvelop

import (
	"fmt"
	"github.com/birjemin/qqenvelop/utils"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// QPayHbURL send envelop url
const qPayHbURL = "https://api.qpay.qq.com/cgi-bin/hongbao/qpay_hb_mch_send.cgi"

// SendQPayHb model
type SendQPayHb struct {
	MchID       string
	MchName     string
	AppSecret   string
	HTTPRequest *utils.HTTPClient
	QqAppID     string // default empty
	NotifyURL   string // default empty
	Charset     int    // default utf8
}

// ParamsSendQPayHb params of envelop
type ParamsSendQPayHb struct {
	TotalAmount int
	Wishing     string
	ActName     string
	IconID      int
	TotalNum    int // default 1
	BannerID    int // default empty
	NotSendMsg  int // default 0
	MinValue    int // default TotalAmount
	MaxValue    int // default TotalAmount
}

// RespSendQPayHb response
type RespSendQPayHb struct {
	RetCode    string `json:"retcode"`
	RetMsg     string `json:"retmsg"`
	ReturnCode string `json:"return_code"`
	ReturnMsg  string `json:"return_msg"`
	ListID     string `json:"listid"`
}

// SendQPayHb send red envelop
// Doc: https://mp.qpay.tenpay.com/buss/wiki/221/1220
func (s *SendQPayHb) SendQPayHb(OpenID string, attributes ParamsSendQPayHb) (*RespSendQPayHb, error) {
	return s.doSendQPayHb(qPayHbURL, OpenID, attributes)
}

// doSendQPayHb handle this action
func (s *SendQPayHb) doSendQPayHb(URL, OpenID string, attributes ParamsSendQPayHb) (*RespSendQPayHb, error) {

	total := strconv.Itoa(attributes.TotalAmount)
	params := map[string]string{
		"charset":      strconv.Itoa(s.getCharset()),
		"nonce_str":    utils.Hex(10),
		"mch_billno":   s.generateBillNo(),
		"mch_id":       s.MchID,
		"mch_name":     s.MchName,
		"re_openid":    OpenID,
		"total_amount": total,
		"total_num":    strconv.Itoa(s.getTotalNum(attributes.TotalNum)),
		"wishing":      attributes.Wishing,
		"act_name":     attributes.ActName,
		"icon_id":      strconv.Itoa(attributes.IconID),
		"not_send_msg": strconv.Itoa(attributes.NotSendMsg),
	}

	// 接收红包者openid或者uin，当re_openid是openid时，qqappid必填
	if s.QqAppID != "" {
		params["qqappid"] = s.QqAppID
	}
	if s.NotifyURL != "" {
		params["notify_url"] = s.NotifyURL
	}
	if attributes.BannerID != 0 {
		params["banner_id"] = strconv.Itoa(attributes.BannerID)
	}
	if attributes.MinValue == 0 {
		params["min_value"] = total
	} else {
		params["min_value"] = strconv.Itoa(attributes.MinValue)
	}

	if attributes.MaxValue == 0 {
		params["max_value"] = total
	} else {
		params["max_value"] = strconv.Itoa(attributes.MaxValue)
	}

	params["sign"] = generateSign(generateQueryStr(s.AppSecret, params))

	var resp = new(RespSendQPayHb)

	if err := s.HTTPRequest.HTTPPost(URL, params); err != nil {
		log.Println("[send]do, post failed", err)
		return resp, err
	}

	if err := s.HTTPRequest.GetResponseJSON(resp); err != nil {
		log.Println("[send]do, response json failed", err)
		return resp, err
	}
	return resp, nil
}

// generateBillNo get bill_no
func (s *SendQPayHb) generateBillNo() string {

	return s.MchID + utils.GetDateNum() + strconv.Itoa(utils.RandNum(1000000000, 10000000000))
}

// getCharset get charset
// option: 1 - utf8, 2 - gbk
// default value: 1
func (s *SendQPayHb) getCharset() int {
	if s.Charset != 2 {
		s.Charset = 1
	}
	return s.Charset
}

// getTotalNum get
// desc: total_num default 1
func (s *SendQPayHb) getTotalNum(num int) int {
	if num <= 0 {
		num = 1
	}
	return num
}

// generateSign generate sign
func generateSign(queryStr string) string {
	return strings.ToUpper(utils.GetMD5String(queryStr))
}

// generateQueryStr generate query str
// doc: https://mp.qpay.tenpay.com/buss/wiki/221/1244
func generateQueryStr(appSecret string, params map[string]string) string {

	query := utils.QuerySortByKeyStr2(params)
	query, _ = url.QueryUnescape(query)
	return fmt.Sprintf("%s&key=%s", query, appSecret)
}
