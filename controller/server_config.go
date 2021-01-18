package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lemon-cloud-project/lemon-cloud-service/service"
	"sync"
)

var ServerConfigControllerInstance *ServerConfigController
var ServerConfigControllerRegisterOnce sync.Once

type ServerConfigController struct {
	serverConfigService service.ServerConfigService
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
