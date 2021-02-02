package main

import (
	"adventure/advserver/config"
	"adventure/advserver/log"
	"adventure/advserver/model"
	"adventure/advserver/msghandler"
	"adventure/advserver/network"
	"adventure/advserver/service"
	"adventure/advserver/sessions"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var logger = log.GetLogger()

func SignalProc() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT /*, syscall.SIGUSR1, syscall.SIGUSR2*/)

	for {
		<-ch
		os.Exit(0)
		return
	}
}

func main() {
	// Init config
	if err := config.Init(); err != nil {
		fmt.Println("config init error")
		os.Exit(-1)
	}
	conf := config.GetConfig()

	// Init log
	if err := log.Init(conf.LogPath, conf.LogLevel); err != nil {
		logger.Error("InitLogger failed (%v)", err)
		os.Exit(-1)
	}

	// Init service
	if err := service.Init(); err != nil {
		fmt.Println("serviece init error")
		return
	}

	// Init model
	if err := model.Init(); err != nil {
		fmt.Println("model init error")
		return
	}

	// Init sessions
	if err := sessions.Init(); err != nil {
		fmt.Println("sesssion message init error")
		return
	}

	// Init msghandler
	if err := msghandler.Init(); err != nil {
		fmt.Println("msghandler init error")
		return
	}

	// Listen to system signal
	go SignalProc()

	// Accept connection
	listener := network.NewTCPServer()
	listener.Run(msghandler.Dispatch)
}
