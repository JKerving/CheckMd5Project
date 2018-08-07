package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var fileName string

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

//创建存储文件绝对路径的记录级文件
func createRecordFile(fileName string) {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(fileName)
			checkError(err)
		}
	} else {
		err := os.Remove(fileName)
		checkError(err)
		_, err = os.Create(fileName)
		checkError(err)
	}
}

//找出所有文件，将文件绝对路径记录至文件中
func listFile(folder string) {
	files, _ := ioutil.ReadDir(folder)
	//判断test.txt文件是否存在
	f, err := os.OpenFile("test.txt", os.O_RDONLY|os.O_APPEND, 0666)
	checkError(err)
	for _, file := range files {
		if file.IsDir() {
			listFile(folder + `\` + file.Name())
		} else {
			fileName = folder + `\` + file.Name()
			checkError(err)
			_, err = f.Write([]byte(fileName + "\r\n"))
			checkError(err)
		}
	}
}

//从记录级文件中按行读取每个文件绝对路径
func readFileByLine(fileName string) []string {
	var str []string
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		str = append(str, line)
		if err != nil {
			if err == io.EOF {
				return str
			}
			return []string{}
		}
	}
	return str
}

//根据每个文件绝对路径，读取文件内容并计算MD5值
func computeMd5ByFilename(fileNameDir []string) map[string]string {
	fileMd5 := make(map[string]string)
	for _, i := range fileNameDir {
		data, err := ioutil.ReadFile(i)
		if err != nil {
			if err == io.EOF {
				panic(err)
			}
		}
		value := md5.Sum(data)
		md5Str := fmt.Sprintf("%x", value)
		fileMd5[i] = md5Str
	}
	return fileMd5
}

func main() {
	svnFolder := `C:\JKerving\Work Documents\培训`
	createRecordFile("test.txt")
	listFile(svnFolder)
	str := readFileByLine("test.txt")
	fmt.Println(len(str))
	f := computeMd5ByFilename(str)
	for k, v := range f {
		fmt.Println(k, v)
	}
}
