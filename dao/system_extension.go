package dao

import (
	"github.com/lemon-cloud-project/lemon-cloud-service/core"
	"github.com/lemon-cloud-project/lemon-cloud-service/entity"
	"gorm.io/gorm"
	"sync"
)

type SystemExtensionDao struct {
}

var systemExtensionDaoInstance *SystemExtensionDao
var systemExtensionDaoInitOnce sync.Once

func SystemExtension() *SystemExtensionDao {
	systemExtensionDaoInitOnce.Do(func() {
		systemExtensionDaoInstance = &SystemExtensionDao{}
	})
	return systemExtensionDaoInstance
}

// 保存系统扩展程序定义信息
func (i *SystemExtensionDao) Save(extension *entity.SystemExtensionEntity) *gorm.DB {
	return core.DB().Save(extension)
}

// 获取所有的系统扩展程序定义
func (i *SystemExtensionDao) ListAll() []entity.SystemExtensionEntity {
	var result []entity.SystemExtensionEntity
	core.DB().Find(&result)
	return result
}
