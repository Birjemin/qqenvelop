package qqenvelop

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/birjemin/qqenvelop/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRet test ret
func TestSend(t *testing.T) {

	ast := assert.New(t)

	var ts = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ret := `{"retcode":"0","retmsg":"ok","listid":"101000000502201506300000100001"}`
		for _, v := range []string{"charset", "nonce_str", "sign", "mch_billno", "total_amount", "wishing"} {
			val := r.FormValue(v)
			if val == "" {
				ret = `{"return_code":"FAIL","return_msg":"","retcode":"066228701","retmsg":"66228701"}`
				break
			}
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(ret)); err != nil {
			t.Fatal(err)
		}
	}))

	certpool := x509.NewCertPool()
	certpool.AddCert(ts.Certificate())

	c := &utils.HTTPClient{
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					ServerName: "example.com",
					RootCAs:    certpool,
				},
			},
		},
	}

	defer ts.Close()

	obj := SendQPayHb{
		MchID:       "",
		MchName:     "",
		AppSecret:   "",
		QqAppID:     "",
		HTTPRequest: c,
	}

	params := ParamsSendQPayHb{
		TotalAmount: 1,
		Wishing:     "新年好",
		ActName:     "新年活动",
		IconID:      23,
	}

	// success
	ret, err := obj.doSendQPayHb(ts.URL, "1", params)
	if err != nil {
		t.Fatal("err: ", err.Error())
	}
	ast.Equal("0", ret.RetCode)

	params = ParamsSendQPayHb{
		TotalAmount: 1,
		Wishing:     "",
		ActName:     "新年活动",
		IconID:      23,
	}

	// failed
	ret, err = obj.doSendQPayHb(ts.URL, "1", params)
	if err != nil {
		t.Fatal("err2: ", err.Error())
	}
	ast.Equal("066228701", ret.RetCode)

}
