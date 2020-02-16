package main

import "RecordLog/myLog"

var mylogger *myLog.LoggerFile

func main() {
	mylogger = myLog.NewFileLog(myLog.DEBUG,"./","test.log")
	defer mylogger.Close()
	for  {
		mylogger.Debug("User with id ten keeps logging in again")
	}
}
