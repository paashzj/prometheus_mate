package main

// 解析Flag
import (
	_ "prometheus_mate/pkg/config"
)

// boot
import (
	_ "prometheus_mate/pkg/boot"
)

// route
import (
	_ "prometheus_mate/pkg/router"
)

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"prometheus_mate/pkg/prom"
)

func main() {
	glog.Info("prom mate started")
	prom.Start()
	g.Server().Run()
}
