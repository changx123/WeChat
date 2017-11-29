package main

import (
	"wechat"
	"fmt"
	"os"
)

func main() {
	var wechat = &wechat.Wechat{}
	wechat.Init()
	uuid , err := wechat.GetUUid()
	if err != nil {
		fmt.Println(err)
		return
	}
	img , err := wechat.GetQrcide(uuid)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.OpenFile("./二维码.png", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write(img)
}
