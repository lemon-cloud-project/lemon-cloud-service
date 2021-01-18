package base

import (
	"gorm.io/gorm"
)

type Entity struct {
	DataKey   string         `gorm:"type:varchar(64);primary_key;comment:'Primary key of this record'" json:"data_key"`
	CreatedAt int64          `gorm:"autoUpdateTime:milli;comment:'Creation time of this record'" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli;comment:'Time of last revision of this record'" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:'The time when this record was soft deleted'" json:"-"`
}
