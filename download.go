package qqenvelop

import (
	"fmt"
	"github.com/birjemin/qqenvelop/utils"
	"strconv"
)

const qPayBillDownload = "https://api.qpay.qq.com/cgi-bin/hongbao/qpay_hb_mch_down_list_file.cgi"

// DownloadQPayHb notify
type DownloadQPayHb struct {
	MchID     string
	AppSecret string
}

// GetDownloadURL get download's url
func (d *DownloadQPayHb) GetDownloadURL(date int) string {

	params := d.GetParams(date)
	query := utils.QuerySortByKeyStr2(params)
	return fmt.Sprintf("%s?%s", qPayBillDownload, query)
}

// GetParams params
func (d *DownloadQPayHb) GetParams(date int) map[string]string {

	params := map[string]string{
		"mch_id": d.MchID,
		"date":   strconv.Itoa(date),
	}
	params["sign"] = generateSign(generateQueryStr(d.AppSecret, params))
	return params
}
