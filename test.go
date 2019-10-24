package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// bingone
func main() {

	// 修改为您的apikey(https://www.yunpian.com)登录官网后获取
	apikey := "81e2d7ae56e5f6ccd66319c9de48195a"
	// 修改为您要发送的手机号码，多个号码用逗号隔开
	mobile := "16606680861"
	// 发送内容
	text := "【云片网】您的验证码是1234"
	// 发送模板编号
	//tpl_id := 1
	// 语音验证码
	//code := "1234"
	//company := "云片网"
	// 发送模板内容
	//tpl_value := url.Values{"#code#": {code}, "#company#": {company}}.Encode()

	// 获取user信息url
	//url_get_user := "https://sms.yunpian.com/v2/user/get.json"
	// 智能模板发送短信url
	url_send_sms := "https://sms.yunpian.com/v2/sms/single_send.json"
	//// 指定模板发送短信url
	//url_tpl_sms := "https://sms.yunpian.com/v2/sms/tpl_single_send.json"
	//// 发送语音短信url
	//url_send_voice := "https://voice.yunpian.com/v2/voice/send.json"

	//data_get_user := url.Values{"apikey": {apikey}}
	data_send_sms := url.Values{"apikey": {apikey}, "mobile": {mobile}, "text": {text}}
	//data_tpl_sms := url.Values{"apikey": {apikey}, "mobile": {mobile},
	//	"tpl_id": {fmt.Sprintf("%d", tpl_id)}, "tpl_value": {tpl_value}}
	//data_send_voice := url.Values{"apikey": {apikey}, "mobile": {mobile}, "code": {code}}

	//httpsPostForm(url_get_user, data_get_user)
	httpsPostForm(url_send_sms, data_send_sms)
	//httpsPostForm(url_tpl_sms, data_tpl_sms)
	//httpsPostForm(url_send_voice, data_send_voice)
}

func httpsPostForm(url string, data url.Values) {
	resp, err := http.PostForm(url, data)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//var item = make(map[string]int)
//
//func main() {
//	//item["name"] = 23
//	//Icon, ok := item["name"]
//	//fmt.Println("icon:", Icon, "ok:", ok)
//	//
//	//queue := make(chan *Request)
//	//go func() {
//	//	i := 0
//	//	for {
//	//		queue <- &Request{id: i}
//	//		time.Sleep(time.Second)
//	//		i++
//	//	}
//	//
//	//}()
//	//
//	////Serve(queue)
//	//Serve1(queue)
//	t := time.Now().Sub(time.Date(2019, 4, 17, 8, 40, 0, 0, time.Local)).Seconds()
//	fmt.Println("float:", t, "int:", int(t))
//	fmt.Println("now_du:", time.Duration(1*60*1e9))
//
//}
//
//type Request struct {
//	id int
//}
//
//var sem = make(chan int, 3)
//
//func handle(r *Request) {
//	sem <- 1   // Wait for active queue to drain.
//	process(r) // May take a long time.
//	<-sem      // Done; enable next request to run.
//}
//
//func Serve(queue chan *Request) {
//	for {
//		req := <-queue
//		fmt.Println("req:", req)
//		go handle(req) // Don't wait for handle to finish.
//	}
//}
//func process(r *Request) {
//	_id := r.id
//	fmt.Printf("这是第%d次的开始, 值为:%d\n", _id, r.id)
//	time.Sleep(time.Second * 2)
//	r.id = 2333
//	fmt.Printf("这是第%d次的结束, 值为:%d\n", _id, r.id)
//}
//
//func Serve1(queue chan *Request) {
//	for req := range queue {
//		sem <- 1
//		go func() {
//			process(req) // Buggy; 见下面的解释。
//			<-sem
//		}()
//		time.Sleep(time.Second * 5)
//		fmt.Println("id:", req.id)
//
//	}
//}
