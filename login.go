package wechat

import (
	"wechat/function"
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
func (w *Wechat) GetQrcide(uuid []byte) ([]byte, error) {
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
