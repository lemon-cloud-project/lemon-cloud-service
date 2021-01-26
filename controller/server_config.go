package controller

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var ServerConfigControllerInstance *ServerConfigController
var ServerConfigControllerRegisterOnce sync.Once

type ServerConfigController struct {
}

func RegisterApiServerConfig(engine *gin.RouterGroup) {
	ServerConfigControllerRegisterOnce.Do(func() {
		ServerConfigControllerInstance = &ServerConfigController{}
		serverConfig := engine.Group("/server_config")
		{
			serverConfig.POST("", ServerConfigControllerInstance.Save)
		}
	})
}

func (i *ServerConfigController) Save(ctx *gin.Context) {

}
