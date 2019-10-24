package main

import (
	"github.com/gin-gonic/gin/json"
	"net/http"
	"path"
)

func main() {
	//http.HandleFunc("/", header)
	//http.HandleFunc("/", text)
	//http.HandleFunc("/", rJson)
	http.HandleFunc("/", file)
	http.ListenAndServe(":3000", nil)

}

//todo 只发送header
func header(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.Header().Set("hahaha", "sleepy")
	w.WriteHeader(201)
}

//TODO 返回普通文本
func text(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

//todo 返回json数据
type Profile struct {
	Name    string
	Hobbies []string
}

func rJson(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//TODO 文件服务

func file(w http.ResponseWriter, r *http.Request) {
	// Assuming you want to serve a photo at 'images/test.png'
	//fp := path.Join("images", "test.png") //output: images/test.png
	fp := path.Join("", "test.png") // output: test.png
	http.ServeFile(w, r, fp)
}
