package prom

import (
	"bufio"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gproc"
	"os"
	"prometheus_mate/pkg/path"
	"prometheus_mate/pkg/prom/generate"
	"prometheus_mate/pkg/util"
)

var ReloadChannel = make(chan struct{})

func Start() {
	startProm()
	go func() {
		for {
			<-ReloadChannel
			startOrReloadProm()
		}
	}()
}

func startOrReloadProm() {
	exists, err := util.ProccessExists(path.PromHome + "/prometheus")
	if err != nil {
		glog.Error("unknown prometheus exists ", err)
		return
	}
	if exists {
		restartProm()
	} else {
		startProm()
	}
}

func startProm() {
	err := generatePromFile()
	if err != nil {
		glog.Error("generate prom config file failed", err)
	}
	gproc.ShellRun("bash -x " + path.PromStartScript)
}

func restartProm() {
	generatePromFile()
	gproc.ShellRun("bash -x " + path.PromReStartScript)
}

func generatePromFile() (err error) {
	file, err := gfile.OpenWithFlag(path.PromConfig, os.O_RDWR|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		glog.Error("open file failed", err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(generate.ConvProm())
	if err != nil {
		glog.Error("write string failed", err)
		return
	}
	err = writer.Flush()
	return
}
