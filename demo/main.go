package main

import (
	"wechat"
	"fmt"
	"os"
)

func main() {
	//获取微信指针
	var we = &wechat.Wechat{}
	//初始化httpx cookie等信息
	we.Init()
	//获取微信uuid
	uuid, err := we.GetUUid()
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取登录二维码
	img, err := we.GetQrcide(uuid)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.OpenFile("./二维码.png", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.Write(img)
	//监听扫描二维码
	for {
		head_img, err := we.ListenQrcode(uuid)
		if err != nil {
			if err == wechat.ErrCodeEq408 {
				fmt.Println("扫码超时重新监听")
				continue
			}
			if err == wechat.ErrCodeEq400 {
				fmt.Println("uuid失效请重新获取")
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Println("base64微信头像:" + string(head_img))
		break
	}
	var redirect_uri []byte
	//监听确认登录
	for {
		redirect_uri, err = we.ConfirmQrcode(uuid)
		if err != nil {
			if err == wechat.ErrCodeEq408 {
				fmt.Println("扫码超时重新监听")
				continue
			}
			if err == wechat.ErrCodeEq400 {
				fmt.Println("uuid失效请重新获取")
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Println("登录成功回调地址:" + string(redirect_uri))
		break
	}
	loginInfo , err := we.GetLoginInfo(redirect_uri)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("登录信息:",loginInfo)
	userInfo , err := we.GetUserInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("用户信息:",userInfo)
}
