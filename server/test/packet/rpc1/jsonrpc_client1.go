package main

//
//import (
//	"fmt"
//	"log"
//	"net/rpc/jsonrpc"
//	"time"
//)
//
//// 算数运算请求结构体
//type ArithRequest struct {
//	A int
//	B int
//}
//
//// 算数运算响应结构体
//type ArithResponse struct {
//	Pro int // 乘积
//	Quo int // 商
//	Rem int // 余数
//}
//
//func main() {
//	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8096")
//	if err != nil {
//		log.Fatalln("dailing error: ", err)
//	}
//
//	req := ArithRequest{9, 2}
//	var res ArithResponse
//
//	//conn.Call() 同步调用
//	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
//	if err != nil {
//		log.Fatalln("arith error: ", err)
//	}
//	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
//
//	//conn.Call() 同步调用
//	err = conn.Call("Arith.Divide", req, &res)
//	if err != nil {
//		log.Fatalln("arith error: ", err)
//	}
//
//	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
//
//	//conn.Go() 异步调用
//	call := conn.Go("Arith.Divide", req, &res, nil)
//	//使用select模型监听通道有数据时执行，否则执行后续程序
//	for {
//		select {
//		case <-call.Done:
//			fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
//		default:
//			fmt.Println("继续向下执行....")
//			time.Sleep(time.Second * 1)
//		}
//	}
//}
