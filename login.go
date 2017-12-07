package wechat

import (
	"wechat/function"
	"encoding/xml"
)

//获取web微信uuid
//return [uuid],[error]
func (w *Wechat) GetUUid() ([]byte, error) {
	//获取时间搓19位纳秒
	sTime := function.GetNewTime()
	//截取0-13位
	sTime = function.RegexpString(sTime, 0, 13)
	//获取uuid
	resp, err := w.httpx.Get("https://login.wx.qq.com/jslogin?appid=wx782c26e4c19acffb&redirect_uri=https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage&fun=new&lang=zh_CN&_=" + sTime)
	if err != nil && resp.Status != 200 {
		return nil, err
	}
	//[window.QRLogin.code = 200; window.QRLogin.uuid = "QbsAGoam5Q==";]
	//获取code 22-25位
	code := string(resp.Body[22:25])
	if code != "200" {
		return nil, ErrCodeNo200
	}
	//获取uuid 50-62位
	uuid := resp.Body[50:62]
	return uuid, nil
}

//获取二维码
//return [img二维码],[error]
func (w *Wechat) GetQrcode(uuid []byte) ([]byte, error) {
	resp, err := w.httpx.Get("https://login.weixin.qq.com/qrcode/" + string(uuid))
	if err != nil && resp.Status != 200 {
		return nil, err
	}
	return resp.Body, nil
}

//监听二维码是否扫描
//return [base64微信头像],[ErrCodeEq408本次请求超时 ErrCodeEq400 uuid失效]
func (w *Wechat) ListenQrcode(uuid []byte) ([]byte, error) {
	//获取时间搓19位纳秒
	sTime := function.GetNewTime()
	//_ 截取0-13位
	sTime = function.RegexpString(sTime, 0, 13)
	//r 截取3-13
	xTime := "-" + function.RegexpString(sTime, 3, 13)
	//获取uuid
	resp, err := w.httpx.Get("https://login.wx.qq.com/cgi-bin/mmwebwx-bin/login?loginicon=true&uuid=" + string(uuid) + "&tip=1&r=" + xTime + "&_=" + sTime)
	if err != nil && resp.Status != 200 {
		return nil, err
	}
	//window.code=201;window.userAvatar = 'data:img/jpg;base64,***';
	//截取code 12-15位
	code := string(resp.Body[12:15])
	//201为成功扫码
	if code == "201" {
		//截取 微信头像base64[37-最后2位]
		b := resp.Body[37:len(resp.Body)-2]
		return b, nil
	}
	//本次请求超时
	if code == "408" {
		return nil, ErrCodeEq408
	}
	//uuid失效
	if code == "400" {
		return nil, ErrCodeEq400
	}
	//未知错误
	return resp.Body, ErrUnknown
}

//监听确认登录
//return [登录成功回调地址],[ErrCodeEq408本次请求超时 ErrCodeEq400 uuid失效]
func (w *Wechat) ConfirmQrcode(uuid []byte) ([]byte, error) {
	//获取时间搓19位纳秒
	sTime := function.GetNewTime()
	//_ 截取0-13位
	sTime = function.RegexpString(sTime, 0, 13)
	//r 截取3-13
	xTime := "-" + function.RegexpString(sTime, 3, 13)
	//获取uuid
	resp, err := w.httpx.Get("https://login.wx.qq.com/cgi-bin/mmwebwx-bin/login?loginicon=true&uuid=" + string(uuid) + "&tip=0&r=" + xTime + "&_=" + sTime)
	if err != nil && resp.Status != 200 {
		return nil, err
	}
	//window.code=200;window.redirect_uri="ht***041";
	//截取code 12-15位
	code := string(resp.Body[12:15])
	if code == "200" {
		//截取 微信头像base64[37-最后2位]
		b := resp.Body[38:len(resp.Body)-2]
		return b, nil
	}
	//本次请求超时
	if code == "408" {
		return nil, ErrCodeEq408
	}
	//uuid失效
	if code == "400" {
		return nil, ErrCodeEq400
	}
	//未知错误
	return resp.Body, ErrUnknown
}

//微信登录信息
type LoginInfo struct {
	//微信skey 每次随机
	Skey string
	//微信sid 每次随机
	Wxsid string
	//微信uin 每个微信号固定值，用来确定身份
	Wxuin string
	//我想pass_ticket 每次随机
	PassTicket string
}

//获取登录信息
func (w *Wechat) GetLoginInfo(redirect_uri []byte) (*LoginInfo, error) {
	version := string(redirect_uri[8:11])
	switch version {
	case "wx2":
		w.apiUrl = string(redirect_uri[0:19])
	case "wx.":
		w.apiUrl = string(redirect_uri[0:18])
	default:
		return nil, ErrVersion
	}
	w.httpx.Get(w.apiUrl)
	resp, err := w.httpx.Get(string(redirect_uri) + "&fun=new&version=v2&lang=zh_CN")
	if err != nil && resp.Status != 200 {
		return nil, err
	}
	//<error><ret>0</ret><message></message><skey>@cry**a8</skey><wxsid>G+92ZsVYiWRNeiPx</wxsid><wxuin>1**9</wxuin><pass_ticket>**</pass_ticket><isgrayscale>1</isgrayscale></error>
	//解析xml结构
	type XmlData struct {
		Message     string `xml:"message"`
		Ret         string `xml:"ret"`
		Skey        string `xml:"skey"`
		Wxsid       string `xml:"wxsid"`
		Wxuin       string `xml:"wxuin"`
		PassTicket  string `xml:"pass_ticket"`
		Isgrayscale string `xml:"isgrayscale"`
	}
	var Uxml XmlData
	//解析xml
	err = xml.Unmarshal(resp.Body, &Uxml)
	if err != nil {
		return nil, err
	}
	//获取成功
	if Uxml.Ret == "0" {
		var loginInfo LoginInfo
		loginInfo.Skey = Uxml.Skey
		loginInfo.Wxsid = Uxml.Wxsid
		loginInfo.Wxuin = Uxml.Wxuin
		loginInfo.PassTicket = Uxml.PassTicket
		w.logininfo = &loginInfo
		return w.logininfo, nil
	}
	//未知错误
	return nil, ErrUnknown
}

//退出
func (w *Wechat) Logout() error {
	resp, err := w.httpx.Post(w.apiUrl+"cgi-bin/mmwebwx-bin/webwxlogout?redirect=1&type=0&skey="+w.logininfo.Skey, []byte("sid="+w.logininfo.Wxsid+"&uin="+w.logininfo.Wxuin))
	if err != nil || resp.Status != 200 {
		return err
	}
	return nil
}
