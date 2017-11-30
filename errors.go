package wechat

import "errors"

var(
	//code返回不为200
	ErrCodeNo200 = errors.New("wechat: Return data code not eq 200")
	//code等于408
	ErrCodeEq408 = errors.New("wechat: Return data code eq 408")
	//code等于400
	ErrCodeEq400 = errors.New("wechat: Return data code eq 400")
	//未知错误
	ErrUnknown = errors.New("wechat: unknown error")
)