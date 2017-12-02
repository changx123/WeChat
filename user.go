package wechat

import (
	"wechat/function"
	"strconv"
	"encoding/json"
)

//获取
func (w *Wechat) GetUserInfo() (*UserInfoJson, error) {
	//获取时间搓19位纳秒
	sTime := function.GetNewTime()
	//_ 截取0-13位
	sTime = function.RegexpString(sTime, 0, 13)
	//获取随机数15位数字
	de := function.RandInt(999999999999999)
	//组合post参数
	d := []byte(`{"BaseRequest": {"Uin": "` + w.logininfo.Wxuin + `","Sid": "` + w.logininfo.Wxsid + `","Skey": "` + w.logininfo.Skey + `","DeviceID": "e` + strconv.Itoa(de) + `"}}`)
	resp, err := w.httpx.Post("https://wx2.qq.com/cgi-bin/mmwebwx-bin/webwxinit?r="+sTime+"&pass_ticket="+w.logininfo.PassTicket, d)
	if err != nil && resp.Status != 200 {
		return nil, err
	}
	//解析数据
	var userInfoJson UserInfoJson
	err = json.Unmarshal(resp.Body, &userInfoJson)
	if err != nil {
		return nil, err
	}
	//判断是否存在错误
	if userInfoJson.BaseResponse.Ret == 0 {
		return &userInfoJson, nil
	}
	return &userInfoJson, ErrUnknown
}
