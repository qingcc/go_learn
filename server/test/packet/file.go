package main

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	//file_info()	//文件信息
	//permission()	//文件权限
	//bufio_with_cache() //带缓存写入(需要flush才能从缓存中将数据写入到磁盘中
	//copy_file() //复制文件
	//read_bytes() //读取字节
	//read_ReadFile() //ReadFile
	//cache_read() //缓存读
	//bufio_scanner() //
	//Zip() //压缩
	//unzip() //解压缩
	//zip_package() //压缩
	//temp() //临时文件夹 &&//临时文件
	//http_download() //http下载
	hash() //http下载
}

//文件信息
func file_info() {
	// 如果文件不存在，则返回错误
	fileInfo, err = os.Stat("test.csv") //os.Stat() 判断文件是否存在
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

//检查读写权限
func permission() {
	// 这个例子测试写权限，如果没有写权限则返回error。
	// 注意文件不存在也会返回error，需要检查error的信息来获取到底是哪个错误导致。
	file, err := os.OpenFile("test.csv", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
		fmt.Println("err:", err)
	}
	info, err := file.Stat()
	fmt.Println("file:", info, "err:", err)
	file.Close()
	// 测试读权限
	file, err = os.OpenFile("test.csv", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		}
	}
	file.Close()
}

//带缓存写入(需要flush才能从缓存中将数据写入到磁盘中
func bufio_with_cache() {
	// 打开文件，只写
	file, err := os.OpenFile("test.csv", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Printf("---------------创建buffered writer 并写数据到缓存-----------------")
	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)
	// 写字节到buffer
	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
	// 写字符串到buffer
	// 也可以使用 WriteRune() 和 WriteByte()
	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered string\n",
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
	//bufferedWriter.Flush()	// TODO bufio 写入缓存, .Flush 时, 才会写入磁盘, 故在 Flush()之前, 文件中不会被写入数据
	log.Printf("----------------检查缓存中的字节数, 并将缓存中的数据写入磁盘----------------")
	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 写内存buffer到硬盘
	bufferedWriter.Flush()
	log.Printf("------------丢弃没有flush的缓存的内容, 清除错误并将缓存传给另外一个writer用--------------------")
	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
	// 当你想将缓存传给另外一个writer时有用
	bufferedWriter.Reset(bufferedWriter)
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
	log.Printf("--------------重新设置缓存的大小, 只可扩容, 不能设置更小的缓存------------------")
	// 重新设置缓存的大小。
	// 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
	// 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
	// 而是writer的原始大小的缓存，默认是4096。
	// 它的功能主要还是为了扩容。
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)
	// resize后检查缓存的大小
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
}

//复制文件
func copy_file() {
	// 打开原始文件
	originalFile, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create("test_copy.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	// 从源中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)
}

//读取字节
func read_bytes() {
	// Open file for reading
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	//read_bytes_Read(file) //read读取 读取最多N个字节
	//read_bytes_ReadFull(file) //read读取 读取正好N个字节
	//read_bytes_ReadAtLeast(file) //read读取 读取至少N个字节
	read_bytes_ReadAll(file) //read读取 读取全部字节

}

//read读取 读取最多N个字节
func read_bytes_Read(file *os.File) {
	// 从文件中读取len(b)字节的文件。
	// 返回0字节意味着读取到文件尾了
	// 读取到文件会返回io.EOF的error
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

//读取正好N个字节
func read_bytes_ReadFull(file *os.File) {
	// file.Read()可以读取一个小文件到大的byte slice中，
	// 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

//读取至少N个字节
func read_bytes_ReadAtLeast(file *os.File) {
	byteSlice := make([]byte, 512)
	minBytes := 8
	// io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
	if err != nil {
		log.Printf("Data read: %s\n", byteSlice)
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

//读取全部字节
func read_bytes_ReadAll(file *os.File) {
	// os.File.Read(), io.ReadFull() 和
	// io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
	// 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
}

func read_ReadFile() {
	// 读取文件到byte slice中
	data, err := ioutil.ReadFile("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)

}

//缓存读
func cache_read() {
	// 打开文件，创建buffered reader
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)
	// 得到字节，当前指针不变
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)
	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)
	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)
	// 读取到分隔符，包含分隔符，返回byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)
	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read string: %s\n", dataString)
	//这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错
}

func bufio_scanner() {
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。
	// 也可以定制一个SplitFunc类型的分隔函数
	scanner.Split(bufio.ScanWords)
	// scan下一个token.
	success := scanner.Scan()
	if success == false {
		// 出现错误或者EOF是返回Error
		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}
	// 得到数据，Bytes() 或者 Text()
	fmt.Println("First word found:", scanner.Text())
	// 再次调用scanner.Scan()发现下一个token
}

func Zip() {
	// 创建一个打包文件
	outFile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// 创建zip writer
	zipWriter := zip.NewWriter(outFile)
	// 往打包文件中写文件。
	// 这里我们使用硬编码的内容，你可以遍历一个文件夹，把文件夹下的文件以及它们的内容写入到这个打包文件中。
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}
	// 下面将要打包的内容写入到打包文件中，依次写入。
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	// 清理
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func unzip() {
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()
	// 遍历打包文件中的每一文件/文件夹
	for _, file := range zipReader.Reader.File {
		// 打包文件中的文件就像普通的一个文件对象一样
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()
		// 指定抽取的文件名。
		// 你可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中。
		// 我们这个例子使用打包文件中相同的文件名。
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)
		// 抽取项目或者创建文件夹
		if file.FileInfo().IsDir() {
			// 创建文件夹并设置同样的权限
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			//抽取正常的文件
			log.Println("Extracting file:", file.Name)
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()
			// 通过io.Copy简洁地复制文件内容
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func zip_package() { //生成压缩文件, 文件为 test2333.csv, 压缩后缀可以为.gz, .rar,tar, 7z等
	//outputFile, err := os.Create("test2333.csv.gz")
	//outputFile, err := os.Create("test2333.csv.rar")
	//outputFile, err := os.Create("test2333.csv.tar")
	outputFile, err := os.Create("test2333.csv.7z")

	if err != nil {
		log.Fatal(err)
	}
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()
	// 当我们写如到gizp writer数据时，它会依次压缩数据并写入到底层的文件中。
	// 我们不必关心它是如何压缩的，还是像普通的writer一样操作即可。
	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Compressed data written to file.")
}

//临时文件夹
func temp() {
	// 在系统临时文件夹中创建一个临时文件夹
	tempDirPath, err := ioutil.TempDir("", "myTempDir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp dir created:", tempDirPath) //临时文件路径
	// 在临时文件夹中创建临时文件
	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp file created:", tempFile.Name())
	// ... 做一些操作 ...
	//写入内容
	tempFile.WriteString("hehehhehehheeheh哈哈哈哈哈哈呵呵呵呵呵")

	// 关闭文件
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	// 删除我们创建的资源
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}
}

//http下载(即http发起请求, 将响应信息写入文件中)
func http_download() {
	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	url := "http://www.devdungeon.com/archive"
	response, err := http.Get(url)
	defer response.Body.Close()
	// 将HTTP response Body中的内容写入到文件
	// Body满足reader接口，因此我们可以使用ioutil.Copy
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}

func hash() { //读取文件内容, 加密
	// 得到文件内容
	data, err := ioutil.ReadFile("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	// 计算Hash
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
}
