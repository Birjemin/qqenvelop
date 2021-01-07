package qqenvelop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestRet test ret
func TestNotify(t *testing.T) {

	ast := assert.New(t)

	notify := NotifyQPayHb{}

	data := `<xml>
	  <mch_id><![CDATA[1104606907]]></mch_id>
	  <mch_billno><![CDATA[29840058602]]></mch_billno>
	  <listid><![CDATA[10000436560988432048]]></listid>
	  <recv_uin><![CDATA[2344546]]></recv_uin>
	  <total_fee><![CDATA[10]]></total_fee>
	  <time_end><![CDATA[20161025094946]]></time_end>
	  <state><![CDATA[1]]></state>
	  <sign><![CDATA[DE4335434F33C065C449E261DCE08BCF]]></sign>
	  <sign_type><![CDATA[MD5]]></sign_type>
	</xml>
	`
	ret, err := notify.Parse([]byte(data))
	if err != nil {
		t.Fatal("parse err: ", err.Error())
	}

	ast.Equal("1104606907", ret.MchID)

	ast.Equal("<xml><return_code>SUCCESS</return_code></xml>", notify.Response("SUCCESS", ""))

	ast.Equal("<xml><return_code>FAIL</return_code><return_msg>ahh</return_msg></xml>", notify.Response("FAIL", "ahh"))
}
