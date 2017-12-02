package function

import (
	"time"
	"math/rand"
)

//获取整数随机数
func RandInt(i int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(i)
}