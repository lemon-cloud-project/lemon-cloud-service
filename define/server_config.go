package define

import "sync"

type ServerConfigDefine struct {
}

var serverConfigDefineInstance *ServerConfigDefine
var serverConfigDefineOnce sync.Once

func ServerConfig() *ServerConfigDefine {
	serverConfigDefineOnce.Do(func() {
		serverConfigDefineInstance = &ServerConfigDefine{}
	})
	return serverConfigDefineInstance
}

const fileName = "config.json"

func (i *ServerConfigDefine) ServerConfigFileName() string {
	return fileName
}
