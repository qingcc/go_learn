module github.com/qingcc/goblog

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/bitly/go-simplejson v0.5.0
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/boombuler/barcode v1.0.0
	github.com/foolin/gin-template v0.0.0-20190415034731-41efedfb393b
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/contrib v0.0.0-20190923054218-35076c1b2bea
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.49.0
	github.com/go-xorm/builder v0.3.4 // indirect
	github.com/go-xorm/core v0.0.0-20180322150003-0177c08cee88
	//github.com/go-xorm/xorm v0.7.9 // indirect
	github.com/go-xorm/xorm v0.7.6
	github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac // indirect
	github.com/golang/net v0.0.0-20180826012351-8a410e7b638d // indirect
	github.com/golang/text v0.3.0 // indirect
	github.com/gorilla/sessions v1.2.0 // indirect
	github.com/gorilla/websocket v1.4.1
	github.com/lib/pq v1.0.0
	github.com/polaris1119/logger v0.0.0-20170422061149-0233d014769e
	github.com/satori/go.uuid v1.2.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/stretchr/testify v1.4.0
	github.com/typa01/go-utils v0.0.0-20181126045345-a86b05b01c1e // indirect
	github.com/ziutek/mymysql v1.5.4
	xorm.io/core v0.7.0
)

replace (
	github.com/qingcc/goblog => ../goblog
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
