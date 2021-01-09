package qqenvelop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDownloadURL test ret
func TestDownloadURL(t *testing.T) {

	ast := assert.New(t)

	notify := DownloadQPayHb{
		AppSecret: "111",
	}

	downloadURL := notify.GetDownloadURL(20210109)

	ast.Equal(
		"https://api.qpay.qq.com/cgi-bin/hongbao/qpay_hb_mch_down_list_file.cgi?date=20210109&mch_id=&sign=AB0BF09C7218B5307C2FBE0A7D688BD7",
		downloadURL,
		)
}

// TestDownloadURL test ret
func TestDownloadParams(t *testing.T) {

	ast := assert.New(t)

	notify := DownloadQPayHb{
		AppSecret: "111",
	}
	params := notify.GetParams(20210109)

	ast.Equal("AB0BF09C7218B5307C2FBE0A7D688BD7", params["sign"])
}
