package setting

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gfm/core/exec"
	"github.com/gfm/core/fo"
	"github.com/gfm/utils/color"
	"github.com/hpcloud/tail"
	"github.com/valyala/fasthttp"
)

func processTask(line []byte) {
	DDSms(fo.GetApi(fo.ConfContent.ConfName), fo.GetMsg(fo.ConfContent.ConfName))
}

func processTaskS(line string) {
	DDSms(fo.GetApi(fo.ConfContent.ConfName), fo.GetMsg(fo.ConfContent.ConfName))
}

//文件监控
func FileMonitoring(filePath string, hookfn func([]byte)) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bufio.NewReaderSize(f, 32768) //默认defaultVufSize=4096
	rd := bufio.NewReader(f)
	f.Seek(1, 2)
	for {
		line, err := rd.ReadBytes('\n')
		//如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(200 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalln(err)
		}
		go hookfn(line)
	}
}

//文件监控2
func FMonitor(filePath string, hookfn func(string)) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tails, err := tail.TailFile(filePath, config)
	if err != nil {
		log.Fatalln(err)
	}

	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			time.Sleep(time.Second)
			continue
		}
		resjs, _ := simplejson.NewJson([]byte(msg.Text))
		res, _ := resjs.Get("msg").String()
		if res == fo.GetMsg(fo.ConfContent.ConfName) {
			go hookfn(res)
		}
	}

}

//钉钉告警fasthttp
func DDSms(apiurl, msg string) error {
	content := `{
	    "msgtype": "text",
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

func Run(model string) {
	switch {
	case model == "debug":
		fmt.Println("Startting send msg！")
		FileMonitoring(fo.GetLog(fo.ConfContent.ConfName), processTask)
	case model == "pro":
		FMonitor(fo.GetLog(fo.ConfContent.ConfName), processTaskS)
	default:
		log.Fatalln("The parameter is error!")
	}
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
	fmt.Println(color.Cyan("   run  --run"), color.White("	       Start up service"))
	fmt.Println(color.Cyan("        --debug"), color.White("       Start up service with debug"))
	fmt.Println(color.Cyan("   version,--version"), color.White("  Gfm Version"))
	fmt.Println(color.Cyan("   help,--help"), color.White("	       Help"))
	fmt.Println("")
	fmt.Println(color.Yellow(" + -------------------------------------------------------------------- +"))
	fmt.Println("")
}
