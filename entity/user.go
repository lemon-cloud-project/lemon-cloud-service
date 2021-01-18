package entity

import (
	"github.com/lemon-cloud-project/lemon-cloud-service/base"
	"time"
)

type UserEntity struct {
	base.Entity
	Number       string    `gorm:"type:varchar(128);unique;comment:'User custom account number'" json:"number"`
	Phone        string    `gorm:"type:varchar(128);unique;comment:'User phone'" json:"phone"`
	Email        string    `gorm:"type:varchar(256);unique;comment:'User email'" json:"email"`
	Password     string    `gorm:"type:varchar(512);comment:'Password hash'" json:"password"`
	PasswordSalt string    `gorm:"type:varchar(64);comment:'Password salt value'" json:"password_salt"`
	Name         string    `gorm:"type:varchar(128);comment:'User real name'" json:"name"`
	NickName     string    `gorm:"type:varchar(256);comment:'User nickname" json:"nick_name"`
	Birthday     time.Time `gorm:"comment:'User birthday" json:"birthday"`
}
