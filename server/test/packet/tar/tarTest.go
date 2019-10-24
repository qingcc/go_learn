package main

/*
*打包和解包的原理和实现
*1、打包实现原理
*	先创建一个文件x.tar，然后向x.tar写入tar头部信息。打开要被tar的文件，
*	向x.tar写入头部信息，然后向x.tar写入文件信息。重复第二步直到所有文件
*	都被写入到x.tar中，关闭x.tar，整个过程就这样完成了
*2、解包实现原理
*	先打开tar文件，然后从这个tar头部中循环读取存储在这个归档文件内的文件头
* 	信息，从这个文件头里读取文件名，以这个文件名创建文件，然后向这个文件里写入数据
*
* 例子: 压缩单个文件
 */

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"blog_go/util"
)

func main() {
	//if err := TarM("logic", "logic.tar.gz"); err != nil {
	//	fmt.Print(err)
	//}
	//
	//if err := UnTar("runtime", "logic.tar.gz"); err != nil {
	//	fmt.Print(err)
	//}

	num := 1
	t := time.Now().Add(time.Duration(num * 6e10))
	fmt.Println("now:", time.Now())
	fmt.Println("tim:", t)
}
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

//region Remark: tar压缩 Author:Qing
func TarM(src, dst string) (err error) {
	//创建文件
	fw, err := os.Create(dst)
	if err != nil {
		return
	}

	defer fw.Close()

	// 将 tar 包使用 gzip 压缩，其实添加压缩功能很简单，
	// 只需要在 fw 和 tw 之前加上一层压缩就行了，和 Linux 的管道的感觉类似
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	//创建 Tar.Writer 结构
	tw := tar.NewWriter(gw)

	defer tw.Close()

	// 下面就该开始处理数据了，这里的思路就是递归处理目录及目录下的所有文件和目录
	// 这里可以自己写个递归来处理，不过 Golang 提供了 filepath.Walk 函数，可以很方便的做这个事情
	// 直接将这个函数的处理结果返回就行，需要传给它一个源文件或目录，它就可以自己去处理
	// 我们就只需要去实现我们自己的 打包逻辑即可，不需要再去路径相关的事情
	return filepath.Walk(src, func(fileName string, fi os.FileInfo, err error) error {
		// 因为这个闭包会返回个 error ，所以先要处理一下这个
		if err != nil {
			return nil
		}

		// 这里就不需要我们自己再 os.Stat 了，它已经做好了，我们直接使用 fi 即可
		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}

		// 这里需要处理下 hdr 中的 Name，因为默认文件的名字是不带路径的，
		// 打包之后所有文件就会堆在一起，这样就破坏了原本的目录结果
		// strings.TrimPrefix 将 fileName 的最左侧的 / 去掉，
		hdr.Name = strings.TrimPrefix(fileName, string(filepath.Separator))

		//写入文件信息
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		//判断文件是否是标准文件, 如果不是就不处理了
		// 如： 目录，这里就只记录了文件信息，不会执行下面的 copy
		if !fi.Mode().IsRegular() {
			return nil
		}

		//打开文件
		fr, err := os.Open(fileName)
		defer fr.Close()
		if err != nil {
			return err
		}

		//copy 文件数据到 tw
		n, err := io.Copy(tw, fr)
		if err != nil {
			return err
		}
		// 记录下过程，这个可以不记录，这个看需要，这样可以看到打包的过程
		log.Printf("成功打包 %s ，共写入了 %d 字节的数据\n", fileName, n)

		return nil
	})

}

//endregion

//region Remark: 解压缩 Author:Qing
func UnTar(dst, src string) (err error) {
	//打开准备解压的文件
	fr, err := os.Open(src)
	if err != nil {
		return
	}
	defer fr.Close()

	//将打开的文件解压
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case hdr == nil:
			continue
		}

		//处理保存的路径
		dstFileDir := filepath.Join(dst, hdr.Name)

		// 根据 header 的 Typeflag 字段，判断文件的类型
		switch hdr.Typeflag {
		case tar.TypeDir: // 如果是目录时候，创建目录
			//判断目录是否存在, 不存在则创建
			if ok, _ := util.DirectoryExists(dstFileDir); !ok {
				util.DirectoryMkdir(dstFileDir)
			}
		case tar.TypeReg: // 如果是文件就写入到磁盘
			// 创建一个可以读写的文件，权限就使用 header 中记录的权限
			// 因为操作系统的 FileMode 是 int32 类型的，hdr 中的是 int64，所以转换下
			file, err := os.OpenFile(dstFileDir, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			n, err := io.Copy(file, tr)
			if err != nil {
				return err
			}
			// 将解压结果输出显示
			fmt.Printf("成功解压： %s , 共处理了 %d 个字符\n", dstFileDir, n)

			// 不要忘记关闭打开的文件，因为它是在 for 循环中，不能使用 defer
			// 如果想使用 defer 就放在一个单独的函数中
			file.Close()
		}
	}
}

//endregion

//region Remark: 测试tar压缩 Author:Qing
func Tar() {
	//准备打包的源文件
	var srcFile = "output1.txt"
	//打包后的文件
	var desFile = srcFile + ".tar"

	//创建需要打包的文件
	fw, err := os.Create(desFile)
	CheckErr(err)
	defer fw.Close()

	// 通过 fw 创建一个 tar.Writer
	tw := tar.NewWriter(fw)

	//关闭tar, 不成功关闭会导致 tar包不完整
	defer func() {
		if err := tw.Close(); err != nil {
			CheckErr(err)
		}
	}()

	//写入文件信息

	// 通过 Stat 获取 FileInfo，然后通过 FileInfoHeader 得到 hdr tar.*Header
	fInfo, err := os.Stat(srcFile)
	CheckErr(err)

	hdr, err := tar.FileInfoHeader(fInfo, "")
	CheckErr(err)

	// 将 tar 的文件信息 hdr 写入到 tw
	err = tw.WriteHeader(hdr)
	CheckErr(err)

	//写入数据
	// 打开准备写入的文件
	fr, err := os.Open(srcFile)
	CheckErr(err)
	defer fr.Close()

	written, err := io.Copy(tw, fr)
	CheckErr(err)
	log.Printf("共写入了 %d 个字符的数据\n", written)

}

//endregion
