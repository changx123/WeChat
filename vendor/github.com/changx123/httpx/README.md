这是一个轻量级的http库，简化 HTTP方面配置和简化对cookie和head的操作
==

##获得httpx
```go
go get github.com/changx123/httpx
```

##获取对象
```go
var hx Httpx
hx = httpx.NewHttpx()
```

##--配置--

##设置http代理
```go
[object].SetProxy("http://127.0.0.1:9090")
```

##设置302重定向次数
```go
//0是不执行302重定向 >0 执行多少次
[object].SetRedirect(0)
```

##是否跳过证书检查
```go
//当遇见证书不受信任导致无法访问开启本项
[object].SetInsecureSkipVerify(true)
```

##设置x509标准证书
```go
if caCrt, err := ioutil.ReadFile("./xxxx.crt"); err == nil {
    //设置证书
    [object].SetCertPoolx509(caCrt)
}
```

##使用cookieJar自动操作保存cookie信息
```go
    [object].SetAutoSaveCookie(true)
```







