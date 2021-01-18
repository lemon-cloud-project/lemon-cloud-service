package dao

import (
	"github.com/lemon-cloud-project/lemon-cloud-service/base"
	"sync"
)

type UserDao struct {
	base.Dao
}

var userDaoInstance *UserDao
var userDaoInitOnce sync.Once

func User() *UserDao {
	userDaoInitOnce.Do(func() {
		userDaoInstance = &UserDao{}
	})
	return userDaoInstance
}
