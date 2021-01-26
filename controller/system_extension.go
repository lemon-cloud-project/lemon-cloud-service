package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"sync"
)

var SystemExtensionControllerInstance *SystemExtensionController
var SystemExtensionControllerRegisterOnce sync.Once

type SystemExtensionController struct {
}

func RegisterApiSystemExtension(engine *gin.RouterGroup) {
	SystemExtensionControllerRegisterOnce.Do(func() {
		SystemExtensionControllerInstance = &SystemExtensionController{}
		SystemExtension := engine.Group("/system_extension")
		{
			SystemExtension.POST("/install", SystemExtensionControllerInstance.Install)
			SystemExtension.POST("/invoke", SystemExtensionControllerInstance.Invoke)
			SystemExtension.POST("/uninstall", SystemExtensionControllerInstance.UnInstall)
		}
	})
}

func (i *SystemExtensionController) Install(ctx *gin.Context) {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("./workspace/ext1", []string{"--runtimePath=hahah", "--runtimePort=12345", "--platformPort=23456"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
	}
	go func() {
		pid.Wait()
		fmt.Println("卧了个槽")
	}()
	fmt.Printf("The process id is %v", pid)
}

func (i *SystemExtensionController) Invoke(ctx *gin.Context) {
}

func (i *SystemExtensionController) UnInstall(ctx *gin.Context) {
}
