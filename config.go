package wechat

//用户浏览器类型(防封设置)
var UserAgent = []string{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60", "Opera/8.0 (Windows NT 5.1; U; en)", "Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.50", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 9.50", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0", "Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11", "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.133 Safari/534.16", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.11 TaoBrowser/2.0 Safari/536.11", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; LBBROWSER)", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E; LBBROWSER)", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; QQBrowser/7.0.3698.400)", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)", "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.84 Safari/535.11 SE 2.X MetaSr 1.0", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SV1; QQDownload 732; .NET4.0C; .NET4.0E; SE 2.X MetaSr 1.0)", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Maxthon/4.4.3.4000 Chrome/30.0.1599.101 Safari/537.36", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 UBrowser/4.0.3214.0 Safari/537.36"}

//用户信息结构
type UserInfoJson struct {
	BaseResponse struct {
		Ret int `json:"Ret"`
		ErrMsg string `json:"ErrMsg"`
	} `json:"BaseResponse"`
	Count int `json:"Count"`
	ContactList []struct {
		Uin int `json:"Uin"`
		UserName string `json:"UserName"`
		NickName string `json:"NickName"`
		HeadImgURL string `json:"HeadImgUrl"`
		ContactFlag int `json:"ContactFlag"`
		MemberCount int `json:"MemberCount"`
		MemberList []interface{} `json:"MemberList"`
		RemarkName string `json:"RemarkName"`
		HideInputBarFlag int `json:"HideInputBarFlag"`
		Sex int `json:"Sex"`
		Signature string `json:"Signature"`
		VerifyFlag int `json:"VerifyFlag"`
		OwnerUin int `json:"OwnerUin"`
		PYInitial string `json:"PYInitial"`
		PYQuanPin string `json:"PYQuanPin"`
		RemarkPYInitial string `json:"RemarkPYInitial"`
		RemarkPYQuanPin string `json:"RemarkPYQuanPin"`
		StarFriend int `json:"StarFriend"`
		AppAccountFlag int `json:"AppAccountFlag"`
		Statues int `json:"Statues"`
		AttrStatus int `json:"AttrStatus"`
		Province string `json:"Province"`
		City string `json:"City"`
		Alias string `json:"Alias"`
		SnsFlag int `json:"SnsFlag"`
		UniFriend int `json:"UniFriend"`
		DisplayName string `json:"DisplayName"`
		ChatRoomID int `json:"ChatRoomId"`
		KeyWord string `json:"KeyWord"`
		EncryChatRoomID string `json:"EncryChatRoomId"`
	} `json:"ContactList"`
	SyncKey struct {
		Count int `json:"Count"`
		List []struct {
			Key int `json:"Key"`
			Val int `json:"Val"`
		} `json:"List"`
	} `json:"SyncKey"`
	User struct {
		Uin int `json:"Uin"`
		UserName string `json:"UserName"`
		NickName string `json:"NickName"`
		HeadImgURL string `json:"HeadImgUrl"`
		RemarkName string `json:"RemarkName"`
		PYInitial string `json:"PYInitial"`
		PYQuanPin string `json:"PYQuanPin"`
		RemarkPYInitial string `json:"RemarkPYInitial"`
		RemarkPYQuanPin string `json:"RemarkPYQuanPin"`
		HideInputBarFlag int `json:"HideInputBarFlag"`
		StarFriend int `json:"StarFriend"`
		Sex int `json:"Sex"`
		Signature string `json:"Signature"`
		AppAccountFlag int `json:"AppAccountFlag"`
		VerifyFlag int `json:"VerifyFlag"`
		ContactFlag int `json:"ContactFlag"`
		WebWxPluginSwitch int `json:"WebWxPluginSwitch"`
		HeadImgFlag int `json:"HeadImgFlag"`
		SnsFlag int `json:"SnsFlag"`
	} `json:"User"`
	ChatSet string `json:"ChatSet"`
	SKey string `json:"SKey"`
	ClientVersion int `json:"ClientVersion"`
	SystemTime int `json:"SystemTime"`
	GrayScale int `json:"GrayScale"`
	InviteStartCount int `json:"InviteStartCount"`
	MPSubscribeMsgCount int `json:"MPSubscribeMsgCount"`
	MPSubscribeMsgList []interface{} `json:"MPSubscribeMsgList"`
	ClickReportInterval int `json:"ClickReportInterval"`
}