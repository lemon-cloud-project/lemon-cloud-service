package define

import "sync"

type AppInfoDefine struct {
}

var appInfoDefineInstance *AppInfoDefine
var appInfoDefineOnce sync.Once

func AppInfo() *AppInfoDefine {
	appInfoDefineOnce.Do(func() {
		appInfoDefineInstance = &AppInfoDefine{}
	})
	return appInfoDefineInstance
}

const bootScreen = `
   __                              ___ _                 _ 
  / /  ___ _ __ ___   ___  _ __   / __\ | ___  _   _  __| |
 / /  / _ \ '_ ' _ \ / _ \| '_ \ / /  | |/ _ \| | | |/ _' |
/ /__|  __/ | | | | | (_) | | | / /___| | (_) | |_| | (_| |
\____/\___|_| |_| |_|\___/|_| |_\____/|_|\___/ \__,_|\__,_|
`

const name = "Lemon Cloud"
const version = "1.0.0"
const dbTablePrefix = "LC_"
const dbColumnPrefix = "LC_"
const apiAddress = ":23385"

func (i *AppInfoDefine) GetBootScreen() string {
	return bootScreen
}

func (i *AppInfoDefine) GetName() string {
	return name
}

func (i *AppInfoDefine) GetVersion() string {
	return version
}

func (i *AppInfoDefine) GetDbTablePrefix() string {
	return dbTablePrefix
}

func (i *AppInfoDefine) GetDbColumnPrefix() string {
	return dbColumnPrefix
}

func (i *AppInfoDefine) GetApiAddress() string {
	return apiAddress
}
