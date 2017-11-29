package wechat

import "errors"

var(
	ErrCodeNo200 = errors.New("wechat: Return data code not eq 200")
)