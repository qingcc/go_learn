package background

import (
	"container/list"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Job struct {
	Id int64
}

var l *list.List

func PostTest(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	job := new(Job)
	job.Id = id
	sign := make(chan int)
	defer close(sign)

	l = list.New()
	l.PushBack(job)

	fmt.Println("入队")
}

func process(job Job) {
	time.Sleep(time.Second)
	fmt.Println("job:", job)

}
