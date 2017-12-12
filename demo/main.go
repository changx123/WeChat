package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"wechat"
	"encoding/base64"
)

func main() {
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()    //获得路由实例
	//注册接口
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/index.html", "./static/index.html")
	router.Static("/js", "./static/js/")
	router.GET("/wechat/login/socket.io",Login)
	////监听端口
	http.ListenAndServe(":8005", router)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Login(c *gin.Context)  {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		DataType, Data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		if DataType == 1 {
			if string(Data) == "1" {
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
				img, err := we.GetQrcode(uuid)
				if err != nil {
					fmt.Println(err)
					return
				}
				conn.WriteMessage(1,Base64Encoding(img))
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
					conn.WriteMessage(1,head_img)
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
					break
				}
				loginInfo, err := we.GetLoginInfo(redirect_uri)
				if err != nil {
					fmt.Println(err)
					return
				}
				conn.WriteMessage(1,[]byte(loginInfo.Wxuin))
			}
		}
	}
}

//base64加密
func Base64Encoding(b []byte) []byte {
	coder := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	buf := make([]byte, coder.EncodedLen(len(b)))
	coder.Encode(buf, b)
	return buf
}
