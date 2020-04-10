package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//// "IO 读写操作"
//func Copy(srcfilename, dstfilename string) (int, error) {
//
//	// 文件打开
//	srcfile, err := os.OpenFile(srcfilename, os.O_RDONLY, os.ModePerm)
//	HandleErr(err)
//	defer srcfile.Close()
//
//	dstfile, err := os.OpenFile(dstfilename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
//	HandleErr(err)
//	defer dstfile.Close()
//
//	temp, err := os.OpenFile(srcfilename+"temp.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
//	defer temp.Close()
//
//	// 临时文件断点
//	temp.Seek(0, 0)
//	bytes := make([]byte, 100, 100)
//	nbyte, err := temp.Read(bytes)
//	cnt, err := strconv.ParseInt(string(bytes[:nbyte]),10,64)
//	fmt.Println(cnt)
//
//	srcfile.Seek(cnt, 0)
//	dstfile.Seek(cnt, 0)
//	data := make([]byte, 1024, 1024)
//	total := 0
//	for {
//		n, err := srcfile.Read(data)
//		if n == 0 || err == io.EOF {
//			fmt.Println(err)
//			//os.Remove(temp)
//			break
//		}
//		fmt.Printf("读取%d byte\n", n)
//		dstfile.Write(data[:n])
//		total += n
//		temp.Seek(0,0)
//		temp.WriteString(strconv.Itoa(total))
//		fmt.Println(total)
//		HandleErr(err)
//		//if total > 100000 {
//		//	panic("断电")
//		//}
//	}
//	return total, nil
//}
//
//func main() {
//
//	srcfilename := "/home/wk/go/src/github.com/iceriverdog/hellogo/Os_pkg/wk.jpeg"
//
//	dstfilename := "/home/wk/go/src/github.com/iceriverdog/hellogo/Os_pkg/dstWK.jpeg"
//
//	n, err := Copy(srcfilename, dstfilename)
//	HandleErr(err)
//	fmt.Printf("总共%d字节", n)
//
//
//}

// bufio
//func main() {
//	srcfilename := "/home/wk/go/src/github.com/iceriverdog/hellogo/Os_pkg/hello"
//	file1, err := os.Open(srcfilename)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer file1.Close()
//
//	//bufio read
//	b1 := bufio.NewReader(file1)
//	p := make([]byte,1024,1024)
//	n, err := b1.Read(p)
//	fmt.Println(n)
//	fmt.Println(string(p[:n]))
//
//	//bufio scanner
//	b2 := bufio.NewReader(os.Stdin)
//	s, err := b2.ReadString('\n')
//	fmt.Println(s)
//
//	//bufio write
//	file2name := "/home/wk/go/src/github.com/iceriverdog/hellogo/Os_pkg/dsthello"
//	b3,err := os.OpenFile(file2name, os.O_CREATE|os.O_WRONLY, os.ModePerm)
//	b4 := bufio.NewWriter(b3)
//	n2, err := b4.WriteString("hello, worild")
//	fmt.Println(n2)
//	b4.Flush()
//	io.EOF
//}

// ioutil
//func NopCloser(r io.Reader) io.ReadCloser
//func ReadAll(r io.Reader) ([]byte, error)
//func ReadDir(dirname string) ([]os.FileInfo, error)
//func ReadFile(filename string) ([]byte, error)
//func TempDir(dir, pattern string) (name string, err error)
//func TempFile(dir, pattern string) (f *os.File, err error)
//func WriteFile(filename string, data []byte, perm os.FileMode) error

// 遍历文件夹
func ListDir(dirName string, level int) {
	fileInfo, err := ioutil.ReadDir(dirName)
	HandleErr(err)
	s := "|--"
	for i := 0; i < level; i++ {
		s += "|--"
	}
	for i := 0; i < len(fileInfo); i++ {
		fmt.Println(s + fileInfo[i].Name())
		if fileInfo[i].IsDir() {
			ListDir(dirName + "/" + fileInfo[i].Name(), level+1)
		}
	}
}
func main() {
	dirName := "/home/wk/go/src/github.com/iceriverdog/hellogo"
	ListDir(dirName, 0)
}
