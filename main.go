package main

import (
	"fmt"
	"github.com/lemon-cloud-project/lemon-cloud-service/core"
	"github.com/lemon-cloud-project/lemon-cloud-service/define"
	"github.com/lemon-cloud-project/lemon-cloud-service/logger"
	"runtime"
)

func main() {
	fmt.Println(define.AppInfo().GetBootScreen())
	fmt.Printf("-- %c[0;0;32m%s - Version: %s - https://www.lemonit.cn%c[0m --\n\n", 0x1B, define.AppInfo().GetName(), define.AppInfo().GetVersion(), 0x1B)
	logger.Info("SYSTEM ARCH: " + runtime.GOARCH)
	logger.Info("SYSTEM OS: " + runtime.GOOS)
	// Initialize database
	if err := core.InitDb(); err != nil {
		logger.Error("Failed to complete database initialization during system startup.", err)
	}
	// Initialize and run REST API
	_ = core.Api().Start(define.AppInfo().GetApiAddress())
}
