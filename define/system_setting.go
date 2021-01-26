package define

import (
	"github.com/lemon-cloud-project/lemon-cloud-service/entity"
	"sync"
)

type SystemSettingDefine struct {
}

var systemSettingDefineInstance *SystemSettingDefine
var systemSettingDefineInitOnce sync.Once

func SystemSetting() *SystemSettingDefine {
	systemSettingDefineInitOnce.Do(func() {
		systemSettingDefineInstance = &SystemSettingDefine{}
	})
	return systemSettingDefineInstance
}

func (i *SystemSettingDefine) GetAllSettingDefineList() []entity.SystemSettingEntity {
	return []entity.SystemSettingEntity{
		entity.SystemSettingEntity{
			Key:       "",
			Value:     "",
			Name:      "",
			Introduce: "",
		},
	}
}
