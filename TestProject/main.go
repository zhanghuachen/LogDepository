package main

import "RecordLog/myLog"

func main() {
	userId := 10


	f := myLog.NewFileLog(myLog.INFO,"./","test.log")
	f.Info("User with id %d keeps logging in again",userId)

	f = myLog.NewFileLog(myLog.DEBUG,"./","test.log")
	f.Debug("User with id %d keeps logging in again",userId)

	f = myLog.NewFileLog(myLog.TRACE,"./","test.log")
	f.Trace("User with id %d keeps logging in again",userId)


	f = myLog.NewFileLog(myLog.WARN,"./","test.log")
	f.Warm("User with id %d keeps logging in again",userId)


	f = myLog.NewFileLog(myLog.ERROR,"./","test.log")
	f.Error("User with id %d keeps logging in again",userId)

	f = myLog.NewFileLog(myLog.FATAL,"./","test.log")
	f.Fatal("User with id %d keeps logging in again",userId)

	c :=myLog.NewConsoleLog(myLog.DEBUG)
	c.Debug("User with id %d keeps logging in again",userId)

	c =myLog.NewConsoleLog(myLog.TRACE)
	c.Trace("User with id %d keeps logging in again",userId)

	c =myLog.NewConsoleLog(myLog.INFO)
	c.Info("User with id %d keeps logging in again",userId)

	c =myLog.NewConsoleLog(myLog.WARN)
	c.Warm("User with id %d keeps logging in again",userId)

	c =myLog.NewConsoleLog(myLog.ERROR)
	c.Error("User with id %d keeps logging in again",userId)

	c =myLog.NewConsoleLog(myLog.FATAL)
	c.Fatal("User with id %d keeps logging in again",userId)


}
