package entity

import "github.com/lemon-cloud-project/lemon-cloud-service/base"

type DomainEntity struct {
	base.Entity
	Key          string `gorm:"unique;type:string;size:64;comment:'The unique readable key of the domain, once created, it cannot be modified.';" json:"key"`
	Name         string `gorm:"type:string;size:128;comment:'Domain name';" json:"name"`
	Introduce    string `gorm:"type:string;size:10240;comment:'The description of the domain, support markdown format';" json:"introduce"`
	StorageType  string `gorm:"type:string;size:128;comment:'The storage type used is associated with the Key of the storage plugin';" json:"storage_type"`
	StorageParam string `gorm:"type:string;size:10240;comment:'Store the parameter JSON data used by the plugin.';" json:"storage_param"`
}
