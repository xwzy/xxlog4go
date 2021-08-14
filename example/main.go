package main

import (
	"errors"
	logger "github.com/xwzy/xxlog4go"

	"time"
)

func main() {
	if err := logger.SetupLogWithConf("./log.json"); err != nil {
		panic(err)
	}
	defer logger.Close()

	var name = "aaaaa"
	other := errors.New("this is a error")

	for {
		logger.Trace("log4go by %s", name)
		logger.Debug("log4go by %s", name)
		logger.Info("log4go by %s", name)
		logger.Warn("log4go by %s", name)
		logger.Error("log4go by %s", name)
		logger.Fatal("log4go by %s", name)

		logger.TraceV(name, other)
		logger.DebugV(name, other)
		logger.InfoV(name, other)
		logger.WarnV(name, other)
		logger.ErrorV(name, other)
		logger.FatalV(name, other)

		time.Sleep(time.Second * 1)
	}
}
