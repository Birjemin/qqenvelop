package qqenvelop

import (
	"github.com/birjemin/qqenvelop/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestDetailHb test ret
func TestDetailHb(t *testing.T) {

	ast := assert.New(t)

	var ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ret := `{"retcode":"0","retmsg":"ok","listid":"101000000502201506300000100001","state":"2","total_num":"1","recv_num":"1","total_amount":"1","recv_amount":"1","recv_details":[{"create_time":"2021-01-09 15:21:02","uin":"111111"}]}`
		for _, v := range []string{"mch_id", "nonce_str", "sign", "listid"} {
			val := r.URL.Query()[v][0]
			if val == "" {
				ret = `{"retcode":"66201001","retmsg":"请求参数错误"}`
				break
			}
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(ret)); err != nil {
			t.Fatal(err)
		}
	}))

	c := &utils.HTTPClient{
		Client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}

	defer ts.Close()

	obj := QPayHbDetail{
		MchID:       "11",
		AppSecret:   "111",
		HTTPRequest: c,
	}

	params := ParamsQPayHbDetail{
		ListID: "101000000502201506300000100001",
	}

	// success
	ret, err := obj.doGetDetail(ts.URL, params)
	if err != nil {
		t.Fatal("err1: ", err.Error())
	}
	ast.Equal("0", ret.RetCode)

	params = ParamsQPayHbDetail{
		ListID: "",
	}

	// failed
	ret, err = obj.doGetDetail(ts.URL, params)
	if err != nil {
		t.Fatal("err2: ", err.Error())
	}
	ast.Equal("66201001", ret.RetCode)

}
