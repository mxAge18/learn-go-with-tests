package v2

import (
	"fmt"
	"net/http"
	"time"
)
// 10s超时时间设置
var tenSecondTimeout = 10 * time.Second

// 比较URL响应速度
func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

// 可配置超时时间的比较较快的URL
func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <- ping(url2):
		return url2, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func () {
		http.Get(url)
		ch <- true
	}()
	return ch
}