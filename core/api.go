package core

import (
	"github.com/gin-gonic/gin"
	"github.com/lemon-cloud-project/lemon-cloud-commons-golang/logger"
	"github.com/lemon-cloud-project/lemon-cloud-service/controller"
	"github.com/lemon-cloud-project/lemon-cloud-service/define"
	"github.com/lemon-cloud-project/lemon-cloud-service/error_code"
	"github.com/lemon-cloud-project/lemon-cloud-service/model"
	"net/http"
	"sync"
)

var ginCoreInstance *ApiCore
var ginCoreInitOnce sync.Once

type ApiCore struct {
	rootEngine          *gin.Engine
	needAuthRouterGroup *gin.RouterGroup
	freeUrlPool         map[string]struct{}
}

func Api() *ApiCore {
	ginCoreInitOnce.Do(func() {
		ginCoreInstance = &ApiCore{}
		ginCoreInstance.Init("/")
	})
	return ginCoreInstance
}

func (i *ApiCore) Init(baseUrl string) {
	i.rootEngine = gin.New()
	i.rootEngine.Use(gin.Logger())
	i.needAuthRouterGroup = i.rootEngine.Group(baseUrl, i.CheckAuthHandler())
	i.RegisterApi()
}

func (i *ApiCore) Start(address string) error {
	return i.rootEngine.Run(address)
}

func (i *ApiCore) RegisterApi() {
	controller.RegisterApiServerConfig(i.needAuthRouterGroup)
	controller.RegisterApiSystemExtension(i.needAuthRouterGroup)
}

func (i *ApiCore) CheckAuthHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqUrlPath := ctx.Request.URL.Path
		if i.checkUrlIsFree(reqUrlPath) {
			logger.Debug("The free api received a network request: " + reqUrlPath)
			ctx.Next()
		} else {
			if i.checkAuthToken(ctx) {
				ctx.Next()
			} else {
				responseAuthError(ctx)
			}
		}
	}
}

func (i *ApiCore) getFreeUrlList() map[string]struct{} {
	if i.freeUrlPool == nil {
		i.freeUrlPool = map[string]struct{}{
			"/": define.Void{},
		}
	}
	return i.freeUrlPool
}

func (i *ApiCore) checkUrlIsFree(url string) bool {
	return false
}

func (i *ApiCore) checkAuthToken(ctx *gin.Context) bool {
	//jwtTokenStr := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
	//return authService.CheckToken(jwtTokenStr)
	return true
}

func responseAuthError(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, model.HttpResponse{
		Success: false,
		Code:    error_code.CommonUnauthorized,
		Data:    nil,
	})
	ctx.Abort()
}
