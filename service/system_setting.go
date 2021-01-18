package service

import (
	"sync"
)

type SystemSettingService struct {
}

var systemSettingInstance *SystemSettingService
var systemSettingInitOnce sync.Once

func SystemSetting() *SystemSettingService {
	systemSettingInitOnce.Do(func() {
		systemSettingInstance = &SystemSettingService{}
	})
	return systemSettingInstance
}

// Complete the missing items in the system settings table
func (i *SystemSettingService) CompletionAllSystemSetting() error {
	return nil
}
