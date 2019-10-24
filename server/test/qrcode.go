package main

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"log"
	"os"
	"time"
	"blog_go/util"
)

func main() {

	path := "./uploads/images/qrcode/" + time.Now().Format("2006/0102/")
	fmt.Println("path:", path)
	util.DirectoryMkdir(path)

	writePng(path+"test.png", "123456", 300)
}

func writePng(filename, base64 string, wide int) (path string) {
	code, err := qr.Encode(base64, qr.L, qr.Unicode)
	// code, err := code39.Encode(base64)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data: ", code.Content())

	if base64 != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, wide, wide)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, code)
	// err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		log.Fatal(err)
	}

	path = filename
	return
}
