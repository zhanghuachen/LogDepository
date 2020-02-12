package myLog

import (
	"path"
	"runtime"
)

func getCallerInfo(skip int)(Filename ,Funcname string,line int){
	pc , filepath , line , ok :=runtime.Caller(skip)
	if !ok {
		return
	}
	Funcname = runtime.FuncForPC(pc).Name()
	Funcname = path.Base(Funcname)
	Filename = path.Base(filepath)
	return
}
