package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

var inputName = flag.String("name", "CHENJIAN", "Input Your Name.")
var inputAge = flag.Int("age", 27, "Input Your Age")
var inputGender = flag.String("gender", "female", "Input Your Gender")
var inputFlagvar int

func Init() {
	flag.IntVar(&inputFlagvar, "flagname", 1234, "Help")
}

func test() {
	Init()
	flag.Parse()
	// func Args() []string
	// Args returns the non-flag command-line arguments.
	// func NArg() int
	// NArg is the number of arguments remaining after flags have been processed.
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("name=", *inputName)
	fmt.Println("age=", *inputAge)
	fmt.Println("gender=", *inputGender)
	fmt.Println("flagname=", inputFlagvar)
}

func main() {
	test_flag()
}

var intervalFlag interval

func _init() {
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func test_flag() {
	_init()
	flag.Parse()
	fmt.Println(intervalFlag)
}

type interval []time.Duration

//实现String接口
func (i *interval) String() string {
	return fmt.Sprintf("%v", *i)
}

//实现Set接口,Set接口决定了如何解析flag的值
func (i *interval) Set(value string) error {
	//此处决定命令行是否可以设置多次-deltaT
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

/*
*input		 go run server/test/packet/flag.go -deltaT 64m,72h,81s
*output:	[1h4m0s 72h0m0s 1m21s]
 */
