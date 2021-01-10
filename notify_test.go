package qqenvelop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestRet test ret
func TestNotify(t *testing.T) {

	ast := assert.New(t)

	notify := NotifyQPayHb{
		AppSecret: "111",
	}

	data := `
	<?xml version="1.0" encoding="UTF-8" ?>
	<xml>
	<appid><![CDATA[1111111111]]></appid>
	<mch_id><![CDATA[1104606907]]></mch_id>
	<openid><![CDATA[JIDJIHUE]]></openid>
	<out_trade_no><![CDATA[10000436560988432048]]></out_trade_no>
	<sign><![CDATA[3B8E2755D32BEE3AF3D45274F03CCB77]]></sign>
	<sign_type><![CDATA[MD5]]></sign_type>
	<state><![CDATA[1]]></state>
	<time_end><![CDATA[2020-01-09 15:21:02]]></time_end>
	<total_fee><![CDATA[1]]></total_fee>
	<transaction_id><![CDATA[29840058602]]></transaction_id>
	</xml>
	`

	ret, err := notify.Parse([]byte(data))
	if err != nil {
		t.Fatal("parse err: ", err.Error())
	}
	b := notify.CheckSign(ret)

	ast.Equal(true, b)

	ast.Equal("1104606907", ret.MchID)

	ast.Equal("<xml><return_code>SUCCESS</return_code></xml>", notify.Response("SUCCESS", ""))

	ast.Equal("<xml><return_code>FAIL</return_code><return_msg>ahh</return_msg></xml>", notify.Response("FAIL", "ahh"))

	data = `
	<?xml version="1.0" encoding="UTF-8" ?>
	<xml>
	<appid><![CDATA[1111111111]]></appid>
	<mch_id><![CDATA[1104606907]]></mch_id>
	<openid><![CDATA[JIDJIHUE]]></openid>
	<out_trade_no><![CDATA[10000436560988432048]]></out_trade_no>
	<sign><![CDATA[F2E3BCE8E6C973DEB2937CFCEAC00CD3]]></sign>
	<sign_type><![CDATA[MD5]]></sign_type>
	<state><![CDATA[1]]></state>
	<time_end><![CDATA[2020-01-09 15:21:02]]></time_end>
	<total_fee><![CDATA[1]]></total_fee>
	<attach><![CDATA[11]]></attach>
	<transaction_id><![CDATA[29840058602]]></transaction_id>
	</xml>
	`

	ret, err = notify.Parse([]byte(data))
	if err != nil {
		t.Fatal("parse err: ", err.Error())
	}
	b = notify.CheckSign(ret)

	ast.Equal(true, b)

	ast.Equal("1104606907", ret.MchID)

	ast.Equal("<xml><return_code>SUCCESS</return_code></xml>", notify.Response("SUCCESS", ""))

	ast.Equal("<xml><return_code>FAIL</return_code><return_msg>ahh</return_msg></xml>", notify.Response("FAIL", "ahh"))
}
