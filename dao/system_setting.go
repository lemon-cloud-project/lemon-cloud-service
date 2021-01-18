package dao

import (
	"github.com/lemon-cloud-project/lemon-cloud-service/core"
	"github.com/lemon-cloud-project/lemon-cloud-service/entity"
	"gorm.io/gorm"
	"sync"
)

type SystemSettingDao struct {
}

var systemSettingDaoInstance *SystemSettingDao
var systemSettingDaoInitOnce sync.Once

func SystemSetting() *SystemSettingDao {
	systemSettingDaoInitOnce.Do(func() {
		systemSettingDaoInstance = &SystemSettingDao{}
	})
	return systemSettingDaoInstance
}

func (i *SystemSettingDao) FindByKey(key string) entity.SystemSettingEntity {
	result := entity.SystemSettingEntity{}
	core.DB().Where(&entity.SystemSettingEntity{
		Key: key,
	}).First(result)
	return result
}

func (i *SystemSettingDao) Save(setting *entity.SystemSettingEntity) *gorm.DB {
	return core.DB().Save(setting)
}
