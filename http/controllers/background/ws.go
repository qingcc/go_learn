package background

import (
	"net/http"
	"net/url"
)

var broadcast = make(chan *SendData)

func init() {
	go wstest()
}

func wstest() {
	for {
		a := <-broadcast
		http.PostForm("http://192.168.20.21:6008/send", url.Values{
			"name": {a.Name},
			"data": {a.Other},
		})
	}
}
