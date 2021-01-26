package main

import (
	"fmt"
	"github.com/lemon-cloud-project/lemon-cloud-commons-golang/logger"
	"github.com/lemon-cloud-project/lemon-cloud-service/core"
	"github.com/lemon-cloud-project/lemon-cloud-service/define"
	"runtime"
)

func main() {
	// 打印基础信息
	fmt.Println(define.AppInfo().GetBootScreen())
	fmt.Printf("-- %c[0;0;32m%s - Version: %s - https://www.lemonit.cn%c[0m --\n\n", 0x1B, define.AppInfo().GetName(), define.AppInfo().GetVersion(), 0x1B)
	logger.Info("SYSTEM ARCH: " + runtime.GOARCH)
	logger.Info("SYSTEM OS: " + runtime.GOOS)
	// 初始化数据库
	if err := core.InitDb(); err != nil {
		logger.Error("Failed to complete database initialization during system startup.", err)
	}
	// 初始化HTTP服务，然后启动
	_ = core.Api().Start(define.AppInfo().GetApiAddress())
}
