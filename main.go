package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"qfw/util"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/p/mahonia"
	"github.com/robfig/cron"
	"github.com/wuxicn/pipeline"
)

//全局配置文件
var Sysconfig map[string]interface{}

//所有定时任务
var jobs = []*program{}

//调度主方法
var CRON *cron.Cron

/**
对于tomcat,一定要设置
1.startup.bat
set CATALINA_HOME=D:\soft\tomcat2
2.同时命名不同的java进程名称：
在java\bin目录下复制出进程名称,如：java3540.exe
3.setclasspath.bat
set _RUNJAVA="%JRE_HOME%\bin\java3540.exe"
**/

func init() {
	ReadConfig(&Sysconfig)
	pros := util.ObjArrToMapArr(Sysconfig["program"].([]interface{}))
	for _, v := range pros {
		tmp := program{}
		tmp.name = v["name"].(string)
		tmp.startcmd = v["startcmd"].(string)
		tmp.findname = v["findname"].(string)
		tmp.cron = v["cron"].(string)
		if tmp.isRun(5) {
			tmp.status = 1
		} else {
			go tmp.start()
		}
		tmp.cronstatus = 1
		jobs = append(jobs, &tmp)
	}
	CRON = cron.New()
}

func main() {
	CRON.Start()
	defer CRON.Stop()
	go run()
	b := make(chan bool)
	<-b
}

func run() {
	log.Println("run cron....")
	for _, p := range jobs {
		pro := p
		log.Println("加载定时任务:", pro.name, pro.cron)
		CRON.AddFunc(pro.cron, func() { pro.restart(3) })
	}
}

type program struct {
	name       string //任务名称
	startcmd   string //启动命令
	findname   string //关闭命令
	cron       string //定时任务
	status     int    //启动状态
	cronstatus int    //定时任务状态
	lasttime   int64  //上次启动时间
	lock       sync.Mutex
}

func (p *program) restart(n int) {
	defer util.Catch()
	log.Println("开始重启：", p.name)
	if p.isRun(5) {
		log.Println("正在关闭：", p.name, "...")
		p.stop()
		time.Sleep(2 * time.Second)
	}
	p.start()
	if p.isRun(10) {
		log.Println("启动成功：", p.name)
	} else {
		log.Println("启动失败：", p.name)
		n--
		if n > 0 {
			p.restart(n)
		}
	}
	return
}

//查询程序是否运行
func (p *program) isRun(n int) (b bool) {
	defer util.Catch()
	for i := 0; i < n; i++ {
		res := execPipe(exec.Command("tasklist"),
			exec.Command("findstr", p.findname),
			exec.Command("find", `/C`, " "))
		intn := util.IntAll(res)
		if intn > 0 {
			b = true
		}
		if b {
			break
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("...checkrun", p.name, b)
	return b
}

func (p *program) stop() {
	defer util.Catch()
	log.Println("...stop", p.name)
	res := execPipe(
		exec.Command("tasklist"),
		exec.Command("findstr", p.findname))
	res = regexp.MustCompile("\\s+").ReplaceAllString(res, " ")

	resArr := strings.Split(res, " ")
	if len(resArr) > 2 && regexp.MustCompile("\\d+").MatchString(resArr[1]) {
		execPipe(exec.Command("tskill", resArr[1]))
	}
}

func (p *program) start() {
	defer util.Catch()
	log.Println("...start", p.name)
	execPipe(exec.Command("cmd.exe", "/c", "call", p.startcmd))
}

//执行管道命令
func execPipe(pipe ...*exec.Cmd) (res string) {
	defer util.Catch()
	stdout, serr, err := pipeline.Run(pipe...)
	log.Println("pipecmd-err-log:", err, string(mahonia.NewDecoder("GBK").ConvertString(serr.String())))
	return strings.TrimRight(string(mahonia.NewDecoder("GBK").ConvertString(stdout.String())), "\r\n")
}

//读取配置文件
func ReadConfig(config ...interface{}) {
	var r *os.File
	filepath := "./config.json"
	if len(config) > 1 {
		filepath, _ = config[0].(string)
	}
	r, _ = os.Open(filepath)
	defer r.Close()
	bs, _ := ioutil.ReadAll(r)
	json.Unmarshal(bs, config[0])
}
