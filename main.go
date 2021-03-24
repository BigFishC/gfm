package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gfm/utils/setting"
)

//一次性读取
func ReadAll(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return ioutil.ReadAll(f)
}

//分块读取，可在速度和内存占用之间取得很好的平衡
func ReadBlock(filePath string, bufSize int, hookfn func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := make([]byte, bufSize) //一次读取多少字节
	/*
		NewReader 创建一个具有默认大小的缓冲区，从r读取的*Reader
	*/
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		hookfn(buf[:n]) //n是成功读取字节数
		if err != nil { //遇到任何错误立即返回，并忽略EOF错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

//输出到控制台

//逐行读取
func ReadLine(filePath string, hookfn func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

//钉钉告警http
func SendMsg(apiurl, msg string) {
	webhook := apiurl
	content := `{"msgtype": "text",
      "text": {"content": "` + msg + `"},
                "at": {
                     "atMobiles": [
                         "18204019490"
                     ],
                     "isAtAll": false
                }
    }`
	//创建一个请求
	req, err := http.NewRequest("POST", webhook, strings.NewReader(content))
	if err != nil {
		fmt.Println("err")
	}
	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "firefox")
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err")
	}
	defer resp.Body.Close()
	// body,_:=ioutil.ReadAll(resp.Body)

}

func main() {

	args := os.Args
	if args == nil || len(args) < 2 {
		setting.Help()
	} else {
		if args[1] == "help" || args[1] == "--help" {
			setting.Help()
		} else if args[1] == "version" || args[1] == "--version" {
			fmt.Println("v0.6.5")
		} else if args[1] == "run" || args[1] == "--run" {
			setting.Run()
		} else {
			setting.Help()
		}
	}
}
