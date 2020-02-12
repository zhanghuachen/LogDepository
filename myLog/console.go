package myLog

import (
	"fmt"
	"os"
	"time"
)

type ConsoleLog struct {
	level Level
}

func NewConsoleLog(level Level) *ConsoleLog{
	con :=  &ConsoleLog{
		level:	level,
	}
	return con
}


func (c *ConsoleLog)log(level Level,format string, args ...interface{}){
	if c.level > level {
		return
	}
	msg := fmt.Sprintf(format,args...)
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	Filename ,Funcname, line  := getCallerInfo(3)
	logmsg := fmt.Sprintf("%s [%s] [%s:%s][%d]%s",nowStr,getLevel(c.level) ,Filename,Funcname,line,msg)
	fmt.Fprintln(os.Stdout, logmsg)

}

func (c *ConsoleLog)Debug(format string, args ...interface{}){
	c.log(DEBUG,format , args...)
}

func (c *ConsoleLog)Trace(format string, args ...interface{}){
	c.log(TRACE,format , args...)
}

func (c *ConsoleLog)Info(format string, args ...interface{}){
	c.log(INFO,format , args...)
}
func (c *ConsoleLog)Error(format string, args ...interface{}){
	c.log(ERROR,format , args...)
}
func (c *ConsoleLog)Fatal(format string, args ...interface{}){
	c.log(FATAL,format , args...)
}
func (c *ConsoleLog)Warm(format string, args ...interface{}){
	c.log(WARN,format , args...)
}

func (c *ConsoleLog)Close()  {

}