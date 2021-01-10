## qqenvelop-qq红包

[![Build Status](https://travis-ci.com/Birjemin/qqenvelop.svg?branch=master)](https://travis-ci.com/Birjemin/qqenvelop) 
[![Go Report Card](https://goreportcard.com/badge/github.com/birjemin/qqenvelop)](https://goreportcard.com/report/github.com/birjemin/qqenvelop) 
[![codecov](https://codecov.io/gh/Birjemin/qqenvelop/branch/master/graph/badge.svg)](https://codecov.io/gh/Birjemin/qqenvelop)


[开发者中心](https://mp.qpay.tenpay.com/buss/wiki/221/1219)

### 引入方式
```
go get github.com/birjemin/qqenvelop
```

### 接口列表

- [创建红包](https://mp.qpay.tenpay.com/buss/wiki/221/1220) ✅
- [领取红包通知（只处理校验）](https://mp.qpay.tenpay.com/buss/wiki/221/1223) ✅
- [对账单下载（只生成下载链接）](https://mp.qpay.tenpay.com/buss/wiki/221/1224) ✅
- [红包详情](https://mp.qpay.tenpay.com/buss/wiki/221/2174) ✅


### 使用方式

```golang

cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
if err != nil {
    panic(err)
}

cTLS := &utils.HTTPClient{
    Client: &http.Client{
        Timeout: 5 * time.Second,
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                Certificates: []tls.Certificate{
                    cert,
                },
            },
        },
    },
}

// 发送红包
sendObj := SendQPayHb{
    MchID:       "商户号",
    MchName:     "商户名称",
    AppSecret:   "app_secret",
    HTTPRequest: cTLS,
}

params := ParamsSendQPayHb{
    TotalAmount: 1,
    Wishing:     "新年好",
    ActName:     "新年活动",
    IconID:      23,
}

// success
ret, err := sendObj.SendQPayHb("open_id", params)
if err != nil {
    panic(err)
}
log.Print("ret: ", ret)

// 查看详情
c := &utils.HTTPClient{
    Client: &http.Client{
        Timeout: 5 * time.Second,
    },
}

obj := QPayHbDetail{
    MchID:       "mch_id",
    AppSecret:   "app_secret",
    HTTPRequest: c,
}

params := ParamsQPayHbDetail{
    ListID: "101000000502201506300000100001",
}

// success
ret, err := obj.GetDetail(params)
if err != nil {
    panic(err)
}
log.Print("ret: ", ret)

// 获取下载链接URL(使用cTLS，或者自定义下载即可)
notify := DownloadQPayHb{
    AppSecret: "app_secret",
}

downloadURL := notify.GetDownloadURL(20210109)

// 处理回调校验
ret, err := notify.Parse(data)
if err != nil {
    panic(err.Error())
}
check := notify.CheckSign(ret)
log.Print("ret: ", check, ret)
```

### 测试
- 测试
    ```
    go test
    ```
- 格式化代码
    ```
    golint
    ```
- 覆盖率
    ```
    go test -cover
    go test -coverprofile=coverage.out 
    go tool cover -html=coverage.out
    ```

### 备注
无