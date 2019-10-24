package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func Save2file(i int, filmname, filmscore, peopleofnum [][]string) {
	f, err := os.Create("C:/Users/Administrator/Desktop/note/测试爬虫/第" + strconv.Itoa(i) + "页.txt")
	if err != nil {
		fmt.Println("os.Create err", err)
		return
	}
	defer f.Close()
	//n:=len(filmname)
	fmt.Println("filmname:", len(filmname))
	fmt.Println("2:", len(filmscore))
	fmt.Println("3:", len(peopleofnum))
	f.WriteString("电影名称" + "\t\t\t" + "电影评分" + "\t\t\t" + "评分人数" + "\r\n")
	for i, v := range filmname {
		//f.WriteString(v[1]+"\t\t\t"+filmscore[i][1]+"\t\t\t"+peopleofnum[i][1]+"\r\n")
		f.WriteString(v[1])

		for j := 0; j < (44 - len(v[1])); j++ {
			f.WriteString(" ")
		}
		f.WriteString("\t")
		f.WriteString(filmscore[i][1])
		for j := 0; j < (44 - len(filmscore[i][1])); j++ {
			f.WriteString(" ")
		}
		f.WriteString("\t")
		f.WriteString(peopleofnum[i][1])
		f.WriteString("\r\n")
		//fmt.Println(len(v[1]))
	}
}
func HttpGetDB(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("http.Get err ", err1)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("文件已读完毕！")
			return
		}
		if err2 != nil && err2 != io.EOF {
			fmt.Println("resp.Body.Read err", err2)
			return
		}
		result += string(buf[:n])
	}
	return
}
func RespTileBD(i int, page chan int) {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="
	result, err := HttpGetDB(url)
	if err != nil {
		fmt.Println("HttpGetDB err ", err)
		return
	}
	ret1 := regexp.MustCompile(`<img width="100" alt="(.*?)" src="`)
	filmname := ret1.FindAllStringSubmatch(result, -1)
	ret2 := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
	filmscore := ret2.FindAllStringSubmatch(result, -1)
	ret3 := regexp.MustCompile(`<span>(.*?)人评价</span>`)
	peopleofnum := ret3.FindAllStringSubmatch(result, -1)
	Save2file(i, filmname, filmscore, peopleofnum)
	page <- i
}
func WorkingDB(start, end int) {
	fmt.Printf("正在爬取%d，到%d页...\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		//rand.Seed(time.Now().UnixNano())
		//n:=rand.Intn(5)+1
		time.Sleep(2 * time.Second)
		go RespTileBD(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("%d页爬取完成\n", <-page)
	}
}
func main() {
	var start, end int
	fmt.Println("请输入起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入终止页：")
	fmt.Scan(&end)
	WorkingDB(start, end)
}
