package wechat

import (
	"wechat/function"
	"strconv"
	"encoding/json"
)

//用户信息结构
type UserInfoJson struct {
	BaseResponse struct {
		Ret    int    `json:"Ret"`
		ErrMsg string `json:"ErrMsg"`
	} `json:"BaseResponse"`
	Count int `json:"Count"`
	ContactList []struct {
		Uin              int           `json:"Uin"`
		UserName         string        `json:"UserName"`
		NickName         string        `json:"NickName"`
		HeadImgURL       string        `json:"HeadImgUrl"`
		ContactFlag      int           `json:"ContactFlag"`
		MemberCount      int           `json:"MemberCount"`
		MemberList       []interface{} `json:"MemberList"`
		RemarkName       string        `json:"RemarkName"`
		HideInputBarFlag int           `json:"HideInputBarFlag"`
		Sex              int           `json:"Sex"`
		Signature        string        `json:"Signature"`
		VerifyFlag       int           `json:"VerifyFlag"`
		OwnerUin         int           `json:"OwnerUin"`
		PYInitial        string        `json:"PYInitial"`
		PYQuanPin        string        `json:"PYQuanPin"`
		RemarkPYInitial  string        `json:"RemarkPYInitial"`
		RemarkPYQuanPin  string        `json:"RemarkPYQuanPin"`
		StarFriend       int           `json:"StarFriend"`
		AppAccountFlag   int           `json:"AppAccountFlag"`
		Statues          int           `json:"Statues"`
		AttrStatus       int           `json:"AttrStatus"`
		Province         string        `json:"Province"`
		City             string        `json:"City"`
		Alias            string        `json:"Alias"`
		SnsFlag          int           `json:"SnsFlag"`
		UniFriend        int           `json:"UniFriend"`
		DisplayName      string        `json:"DisplayName"`
		ChatRoomID       int           `json:"ChatRoomId"`
		KeyWord          string        `json:"KeyWord"`
		EncryChatRoomID  string        `json:"EncryChatRoomId"`
	} `json:"ContactList"`
	SyncKey struct {
		Count int `json:"Count"`
		List []struct {
			Key int `json:"Key"`
			Val int `json:"Val"`
		} `json:"List"`
	} `json:"SyncKey"`
	User struct {
		Uin               int    `json:"Uin"`
		UserName          string `json:"UserName"`
		NickName          string `json:"NickName"`
		HeadImgURL        string `json:"HeadImgUrl"`
		RemarkName        string `json:"RemarkName"`
		PYInitial         string `json:"PYInitial"`
		PYQuanPin         string `json:"PYQuanPin"`
		RemarkPYInitial   string `json:"RemarkPYInitial"`
		RemarkPYQuanPin   string `json:"RemarkPYQuanPin"`
		HideInputBarFlag  int    `json:"HideInputBarFlag"`
		StarFriend        int    `json:"StarFriend"`
		Sex               int    `json:"Sex"`
		Signature         string `json:"Signature"`
		AppAccountFlag    int    `json:"AppAccountFlag"`
		VerifyFlag        int    `json:"VerifyFlag"`
		ContactFlag       int    `json:"ContactFlag"`
		WebWxPluginSwitch int    `json:"WebWxPluginSwitch"`
		HeadImgFlag       int    `json:"HeadImgFlag"`
		SnsFlag           int    `json:"SnsFlag"`
	} `json:"User"`
	ChatSet             string        `json:"ChatSet"`
	SKey                string        `json:"SKey"`
	ClientVersion       int           `json:"ClientVersion"`
	SystemTime          int           `json:"SystemTime"`
	GrayScale           int           `json:"GrayScale"`
	InviteStartCount    int           `json:"InviteStartCount"`
	MPSubscribeMsgCount int           `json:"MPSubscribeMsgCount"`
	MPSubscribeMsgList  []interface{} `json:"MPSubscribeMsgList"`
	ClickReportInterval int           `json:"ClickReportInterval"`
}

//获取用户数据
func (w *Wechat) GetUserInfo() (*UserInfoJson, error) {
	//获取时间搓19位纳秒
	sTime := function.GetNewTime()
	//_ 截取0-13位
	sTime = function.RegexpString(sTime, 0, 13)
	//r 截取3-13
	xTime := "-" + function.RegexpString(sTime, 3, 13)
	//获取随机数15位数字
	de := function.RandInt(999999999999999)
	//组合post参数
	d := []byte(`{"BaseRequest": {"Uin": "` + w.logininfo.Wxuin + `","Sid": "` + w.logininfo.Wxsid + `","Skey": "` + w.logininfo.Skey + `","DeviceID": "e` + strconv.Itoa(de) + `"}}`)
	resp, err := w.httpx.Post(w.apiUrl+"cgi-bin/mmwebwx-bin/webwxinit?r="+xTime+"&pass_ticket="+w.logininfo.PassTicket, d)
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

//用户好友列表结构
type UserFriendListJson struct {
	BaseResponse struct {
		ErrMsg string `json:"ErrMsg"`
		Ret    int    `json:"Ret"`
	} `json:"BaseResponse"`
	MemberCount int `json:"MemberCount"`
	MemberList []struct {
		Alias            string        `json:"Alias"`
		AppAccountFlag   int           `json:"AppAccountFlag"`
		AttrStatus       int           `json:"AttrStatus"`
		ChatRoomID       int           `json:"ChatRoomId"`
		City             string        `json:"City"`
		ContactFlag      int           `json:"ContactFlag"`
		DisplayName      string        `json:"DisplayName"`
		EncryChatRoomID  string        `json:"EncryChatRoomId"`
		HeadImgURL       string        `json:"HeadImgUrl"`
		HideInputBarFlag int           `json:"HideInputBarFlag"`
		IsOwner          int           `json:"IsOwner"`
		KeyWord          string        `json:"KeyWord"`
		MemberCount      int           `json:"MemberCount"`
		MemberList       []interface{} `json:"MemberList"`
		NickName         string        `json:"NickName"`
		OwnerUin         int           `json:"OwnerUin"`
		PYInitial        string        `json:"PYInitial"`
		PYQuanPin        string        `json:"PYQuanPin"`
		Province         string        `json:"Province"`
		RemarkName       string        `json:"RemarkName"`
		RemarkPYInitial  string        `json:"RemarkPYInitial"`
		RemarkPYQuanPin  string        `json:"RemarkPYQuanPin"`
		Sex              int           `json:"Sex"`
		Signature        string        `json:"Signature"`
		SnsFlag          int           `json:"SnsFlag"`
		StarFriend       int           `json:"StarFriend"`
		Statues          int           `json:"Statues"`
		Uin              int           `json:"Uin"`
		UniFriend        int           `json:"UniFriend"`
		UserName         string        `json:"UserName"`
		VerifyFlag       int           `json:"VerifyFlag"`
	} `json:"MemberList"`
	Seq int `json:"Seq"`
}

//获取用户好友列表
func (w *Wechat) GetUserFriend() (*UserFriendListJson, error) {
	//获取时间搓19位纳秒
	sTime := function.GetNewTime()
	//_ 截取0-13位
	sTime = function.RegexpString(sTime, 0, 13)
	//r 截取3-13
	xTime := "-" + function.RegexpString(sTime, 3, 13)
	//请求数据
	resp, err := w.httpx.Get(w.apiUrl + "cgi-bin/mmwebwx-bin/webwxgetcontact?lang=zh_CN&pass_ticket=" + w.logininfo.PassTicket + "&r=" + xTime + "&seq=0&skey=" + w.logininfo.Skey)
	if err != nil && resp.Status != 200 {
		return nil, err
	}
	//解析数据
	var userFriendListJson UserFriendListJson
	err = json.Unmarshal(resp.Body, &userFriendListJson)
	if err != nil {
		return nil, err
	}
	//判断是否存在错误
	if userFriendListJson.BaseResponse.Ret == 0 {
		return &userFriendListJson, nil
	}
	return &userFriendListJson, ErrUnknown
}
