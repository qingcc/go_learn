package main

import (
	"fmt"
	"net/http"
)

func main() {
	//test_nil()
	test_nil_map()
}

func test_nil() {
	t := &tree{v: 10, l: &tree{v: 3}, r: &tree{v: 8}}
	fmt.Println(t.Sum())
	fmt.Println(t.Find(6))
}

type tree struct {
	v int
	l *tree
	r *tree
}

func (t *tree) Sum() int { //二叉树求和
	if t == nil {
		return 0
	}
	return t.v + t.l.Sum() + t.r.Sum()
}

func (t *tree) Find(v int) bool { //二叉树查找某个值
	if t == nil {
		return false
	}
	return t.v == v || t.l.Find(v) || t.r.Find(v)
}

func test_nil_map() {
	//itemMap := make(map[string]string)
	req, err := NewGet("www.google.com", map[string]string{"USER_AGENT": "golang/gopher"})
	//req, err := NewGet("www.google.com", map[string]string{})
	//req, err := NewGet("www.google.com", itemMap)
	fmt.Println("req:", req)
	fmt.Println("err:", err)
}

func NewGet(url string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	headers["ccc"] = "ggg"
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}
