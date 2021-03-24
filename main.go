package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
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
func processTask(line []byte) {
	// os.Stdout.Write(line)
	if string(line[0]) == "3" {
		DDSms("https://oapi.dingtalk.com/robot/send?access_token=bb7e54b59548045909b5042f90dd2e635f56ee9055a3b7e90cbb88821a413536", string(line[0]))
	}
}

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

//文件监控
func FileMonitoring(filePath string, hookfn func([]byte)) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	f.Seek(0, 2)
	for {
		line, err := rd.ReadBytes('\n')
		//如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalln(err)
		}
		go hookfn(line)
	}
}

//钉钉告警
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

func DDSms(apiurl, msg string) error {
	content := `{"msgtype": "text",
      "text": {"content": "` + msg + `"},
                "at": {
                     "atMobiles": [
                         "18204019490"
                     ],
                     "isAtAll": false
                }
    }`

	req := &fasthttp.Request{}
	req.SetRequestURI(apiurl)

	reqBody := []byte(content)
	req.SetBody(reqBody)

	//设置请求头
	req.Header.SetContentType("application/json")

	//设置请求方式
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{}
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return err
	}
	return nil
}

func main() {
	//一次性读取1
	//直接读取文件，无需打开句柄

	// ret, err := ioutil.ReadFile("/Users/tony/text.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(ret)

	//一次性读取
	// ret, err := ReadAll("/Users/tony/text.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(ret)

	//分块读取
	// ReadBlock("/Users/tony/text.txt", 10000, processTask)

	//逐行读取
	// ReadLine("/Users/tony/text.txt", processTask)

	//日志监控
	FileMonitoring("C:\\Users\\acer\\Documents\\1.txt", processTask)
	// SendMsg("https://oapi.dingtalk.com/robot/send?access_token=bb7e54b59548045909b5042f90dd2e635f56ee9055a3b7e90cbb88821a413536", "测试")

}
