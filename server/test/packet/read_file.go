package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	//Ioutil("./output1.txt")
	//OsIoutil("./output1.txt")
	//FileRead("./output1.txt")
	BufioRead("./output1.txt")
}

func Ioutil(name string) {
	if contents, err := ioutil.ReadFile(name); err == nil {
		fmt.Println(string(contents))
	}
}

func OsIoutil(name string) {
	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		if contents, err := ioutil.ReadAll(fileObj); err == nil {
			fmt.Println("Use os.Open family functions and ioutil.ReadAll to read a file contents:", string(contents))
		}

	}
}

//在定义空的byte列表时尽量大一些，否则这种方式读取内容可能造成文件读取不完整
func FileRead(name string) {
	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		//在定义空的byte列表时尽量大一些，否则这种方式读取内容可能造成文件读取不完整
		buf := make([]byte, 1024)
		if n, err := fileObj.Read(buf); err == nil {
			fmt.Println("The number of bytes read:"+strconv.Itoa(n), "Buf length:"+strconv.Itoa(len(buf)))
			//result := strings.Replace(string(buf), "\n", "", 1)
			fmt.Println("Use os.Open and File's Read method to read a file:", string(buf))
		}
	}
}

func BufioRead(name string) {
	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		//一个文件对象本身是实现了io.Reader的 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
		reader := bufio.NewReader(fileObj)

		read(reader) //err: <nil> 	\n	n: 17 buf: 测试n	\n 		测试n@
		//discarded(reader)		//err: <nil> 	\n	buf: 8	\n 	使用ReadSlince相关方法读取内容: 	\n	测试n@
		//peak(reader)			//测试n 	测试
		//readSlice(reader)		//line: 测试n err: <nil>
		//readBytes(reader)		//测试n 	测试n@
		//readString(reader)	//测 size: 3
		//readRune(reader)		//测 size: 3
	}
}

//把Reader缓存对象中的数据读入到[]byte类型的p中，并返回读取的字节数。读取成功，err将返回空值
func read(reader *bufio.Reader) {
	//使用ReadString(delim byte)来读取delim以及之前的数据并返回相关的字符串.
	if result, err := reader.ReadString(byte('@')); err == nil {
		fmt.Println("使用ReadSlince相关方法读取内容:", result)
	}
	//注意:上述ReadString已经将buffer中的数据读取出来了，下面将不会输出内容
	//需要注意的是，因为是将文件内容读取到[]byte中，因此需要对大小进行一定的把控
	buf := make([]byte, 1024)
	//读取Reader对象中的内容到[]byte类型的buf中
	if n, err := reader.Read(buf); err == nil {
		fmt.Println("The number of bytes read:" + strconv.Itoa(n))
		//这里的buf是一个[]byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
		fmt.Println("Use bufio.NewReader and os.Open read file contents to a []byte:", string(buf))
	}
}

//Discard方法跳过后续的 n 个字节的数据，返回跳过的字节数。如果0 <= n <= b.Buffered(),该方法将不会从io.Reader中成功读取数据。
func discarded(reader *bufio.Reader) {
	discarded := 0
	discarded, err := reader.Discard(8)
	fmt.Println("err:", err)
	fmt.Println("buf:", discarded)
	//使用ReadString(delim byte)来读取delim以及之前的数据并返回相关的字符串.
	if result, err := reader.ReadString(byte('@')); err == nil {
		fmt.Println("使用ReadSlince相关方法读取内容:", result)
	}
}

//Peekf方法返回缓存的一个切片，该切片只包含缓存中的前n个字节的数据
func peak(reader *bufio.Reader) {
	buf, err := reader.Peek(15)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf))
}

//该方法在b中读取delimz之前的所有数据，返回的切片是已读出的数据的引用，切片中的数据在下一次的读取操作之前是有效的。
// 如果未找到delim，将返回查找结果并返回nil空值。因为缓存的数据可能被下一次的读写操作修改，因此一般使用ReadBytes或者ReadString，他们返回的都是数据拷贝
func readSlice(reader *bufio.Reader) {
	line, err := reader.ReadSlice('n')
	fmt.Println("line:", string(line), "err:", err)
}

//返回单个字节，如果没有数据返回err
func readByte(reader *bufio.Reader) {
	buf, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf))
}

//返回单个字节，如果没有数据返回err
func readBytes(reader *bufio.Reader) {
	buf, err := reader.ReadBytes('@')
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf))
}

//功能同ReadBytes,返回字符串
func readString(reader *bufio.Reader) {
	buf, err := reader.ReadString('n')
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf))
}

//读取单个UTF-8字符并返回一个rune和字节大小
func readRune(reader *bufio.Reader) {
	buf, size, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf), "size:", size)
}
