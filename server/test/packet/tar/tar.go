package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 创建一个缓冲区来写入我们的存档。
	buf := new(bytes.Buffer)

	// 创建一个新的tar存档。
	tw := tar.NewWriter(buf)

	// 将一些文件添加到存档中。
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}
	if err := tw.Flush(); err != nil {
		log.Fatalln(err)
	}
	// 确保在Close时检查错误。
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	// 打开tar档案以供阅读。
	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)

	// 迭代档案中的文件。
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// tar归档结束
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}

}
