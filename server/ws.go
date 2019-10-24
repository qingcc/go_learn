package main

import "routers"

func main() {
	//ws_ip := logic.ConfigLogic{}.ReadConfig("ws_ip")
	//stp := strings.Split(ws_ip, ":")
	//ip := ":" + stp[len(stp)-1]
	//fmt.Println("ip:", ip)
	router := routers.InitWsRouter()
	//
	router.Run(":6008")
}
