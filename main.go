package main

import (
	"edy-robot/robot"
	llog "github.com/name5566/leaf/log"
	"log"
	"os"
	"os/signal"
)

func init() {
	logger, err := llog.New("debug", "", log.Lshortfile|log.LstdFlags)
	if err != nil {
		panic(err)
	}
	llog.Export(logger)
}

func main() {
	robot.Init()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case sig := <-c:
		llog.Release("closing down (signal: %v)", sig)
		robot.Destroy()
	}
}
