package entity

import "github.com/lemon-cloud-project/lemon-cloud-service/base"

type SystemSettingEntity struct {
	base.Entity
	Key       string `gorm:"type:varchar(128);unique;<-:create;comment:'System setting item key'" json:"key"`
	Value     string `gorm:"type:varchar(10240);<-;comment:'System setting item value'" json:"value"`
	Name      string `gorm:"type:varchar(128);<-:create;comment:'System setting item name'" json:"name"`
	Introduce string `gorm:"type:varchar(1024);<-:create;comment:'System setting item introduce text'" json:"introduce"`
}
