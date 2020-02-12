package myLog

import (
	"fmt"
	"os"
	"path"
	"time"
)

type LoggerFile struct {
	FileName string
	FilePath string
	level 	Level
	FileLog *os.File
	ErrFileLog *os.File
	maxSize int64
}

func NewFileLog(level Level, FilePath, FileName string) *LoggerFile{
	fileObj :=  &LoggerFile{
		level:	level,
		FileName:   FileName,
		FilePath:   FilePath,
		maxSize:	1 * 1024 * 1024,
	}
	fileObj.iniFileLog()
	return fileObj
}
func (f *LoggerFile)iniFileLog(){
	filepath := path.Join(f.FilePath,f.FileName)
	file,err := os.OpenFile(filepath,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		panic(fmt.Errorf("OpenFile err:%v",err))
	}
	f.FileLog = file

	errfile,err := os.OpenFile(filepath,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		panic(fmt.Errorf("OpenFile err:%v",err))
	}
	f.ErrFileLog = errfile
}

func (f *LoggerFile)log(level Level,format string, args ...interface{}){
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format,args...)
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	Filename ,Funcname, line  := getCallerInfo(3)
	msg = fmt.Sprintf("%s [%s] [%s:%s][%d]%s",nowStr,getLevel(level) ,Filename,Funcname,line,msg)
	if f.checkSplit(f.FileLog) {
		f.FileLog = f.splitLogFile(f.FileLog)
	}
	fmt.Fprintln(f.FileLog, msg)
	if level > ERROR {
		if f.checkSplit(f.ErrFileLog) {
			f.ErrFileLog = f.splitLogFile(f.ErrFileLog)
		}
		fmt.Fprintln(f.ErrFileLog,msg)
	}
}

func (f *LoggerFile)checkSplit(file *os.File)bool{
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize
}

func (f *LoggerFile)splitLogFile(file *os.File)*os.File{
		fileName := file.Name()
		backName := fmt.Sprintf("%s_%v.back",fileName,time.Now().Unix())
		file.Close()
		os.Rename(fileName,backName)
		fileObj,err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
		if err != nil {
			panic(fmt.Errorf("OpenFile err:%v",err))
		}
		return fileObj
}

func (f *LoggerFile)Debug(format string, args ...interface{}){
	f.log(DEBUG,format , args...)
}

func (f *LoggerFile)Trace(format string, args ...interface{}){
	f.log(TRACE,format , args...)
}

func (f *LoggerFile)Info(format string, args ...interface{}){
	f.log(INFO,format , args...)
}
func (f *LoggerFile)Error(format string, args ...interface{}){
	f.log(ERROR,format , args...)
}
func (f *LoggerFile)Fatal(format string, args ...interface{}){
	f.log(FATAL,format , args...)
}
func (f *LoggerFile)Warm(format string, args ...interface{}){
	f.log(WARN,format , args...)
}

func (f *LoggerFile)Close(){
	f.FileLog.Close()
	f.ErrFileLog.Close()
}

