package setting

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gfm/core/exec"
	"github.com/gfm/utils/color"
	"github.com/valyala/fasthttp"
)

func processTask(line []byte) {
	// os.Stdout.Write(line)
	if string(line[0]) == "3" {
		DDSms("https://oapi.dingtalk.com/robot/send?access_token=bb7e54b59548045909b5042f90dd2e635f56ee9055a3b7e90cbb88821a413536", string(line[0]))
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

//钉钉告警fasthttp
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

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) //释放使用过的资源

	//设置请求头
	req.Header.SetContentType("application/json")

	//设置请求方式
	req.Header.SetMethod("POST")

	req.SetRequestURI(apiurl)
	reqBody := []byte(content)
	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) //释放使用过的资源

	if err := fasthttp.Do(req, resp); err != nil {
		return err
	}
	return nil
}

func Run() {
	//一次性读取
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

func Help() {
	exec.Execute("clear")
	logo := ` o
 |/\|
 ( OO)                    \|/
 ( \/)  .===O- ~~~biu~biu~ -O-O-
 /   \_/U'                /|\
 ||  |_/
 \\  |	     ~ By: Devops of metro
 {K} ||         _______   _____   ___   ___
  | PP         / ____ /  / ___/  /   | /   |  
  | ||        / /___—/— / ___/  / /| |/ /| |
  (__\\      /______/  /_/     /_/ |_/_/ |_| v0.1.0
`
	fmt.Println(color.Yellow(logo))
	fmt.Println(color.White(" A Log Monitor Software for Metro Service"))
	fmt.Println("")
	fmt.Println(color.Yellow(" + [ ABOUT ]----------------------------------------------------------- +"))
	fmt.Println("")
	fmt.Println(color.Green("   - Github:"), color.White("https://github.com/BigFishC/gfm.git"), color.Green(" - Team:"), color.White("https://navi.nsmetro.cn/navi"))
	fmt.Println("")
	fmt.Println(color.Yellow(" + [ ARGUMENTS ]------------------------------------------------------- +"))
	fmt.Println("")
	fmt.Println(color.Cyan("   run,--run"), color.White("	       Start up service"))
	//fmt.Println(color.Cyan("   init,--init"), color.White("		   Initialization, Wipe data"))
	fmt.Println(color.Cyan("   version,--version"), color.White("  Gfm Version"))
	fmt.Println(color.Cyan("   help,--help"), color.White("	       Help"))
	fmt.Println("")
	fmt.Println(color.Yellow(" + -------------------------------------------------------------------- +"))
	fmt.Println("")
}
