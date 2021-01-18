package core

import (
	"errors"
	"github.com/lemon-cloud-project/lemon-cloud-service/define"
	"github.com/lemon-cloud-project/lemon-cloud-service/entity"
	"github.com/lemon-cloud-project/lemon-cloud-service/logger"
	"github.com/lemon-cloud-project/lemon-cloud-service/model"
	"github.com/lemon-cloud-project/lemon-cloud-service/service"
	"github.com/lemon-cloud-project/lemon-cloud-service/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"reflect"
)

var dbInstance *gorm.DB

func DB() *gorm.DB {
	if dbInstance == nil {
		err := InitDb()
		if err != nil {
			logger.Error("There was an error initializing the database", err)
		}
	}
	return dbInstance
}

func InitDb() error {
	logger.Info("The system started trying to connect to the database.")
	databaseUrl := service.ServerConfig().GetMySQLConnectionUrl()
	if databaseUrl == "" {
		return errors.New("Unable to init the database because the server config is not configured correctly")
	}
	dbObject, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       databaseUrl,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: &model.CustomDbNamingStrategy{
			TablePrefix:   define.AppInfo().GetDbTablePrefix(),
			ColumnPrefix:  define.AppInfo().GetDbColumnPrefix(),
			SingularTable: true,
			NameReplacer:  nil,
		},
	})
	if err != nil {
		return err
	}
	dbInstance = dbObject
	err = dbInstance.Callback().Create().Register("create_generate_data_key", func(db *gorm.DB) {
		if db.Statement.Schema != nil {
			for _, field := range db.Statement.Schema.PrimaryFields {
				if field.Name == "DataKey" {
					switch db.Statement.ReflectValue.Kind() {
					case reflect.Slice, reflect.Array:
						for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
							// 如果DataKey为空，那么生成一个UUID作为DataKey补全上
							fixDataKeyField(db, field, db.Statement.ReflectValue.Index(i))
						}
					case reflect.Struct:
						fixDataKeyField(db, field, db.Statement.ReflectValue)
					}
				}
			}
		}
	})
	dbInstance.Set("gorm:auto_preload", true)
	if err != nil {
		logger.Error("The system could not continue to run because it could not establish a connection with the database.", err)
		logger.Warn("Database connection URL:" + databaseUrl)
	} else {
		updateDb(dbInstance)
		logger.Info("Database connection completed!")
	}
	return err
}

func fixDataKeyField(db *gorm.DB, field *schema.Field, value reflect.Value) {
	if fieldValue, isZero := field.ValueOf(value); !isZero {
		if fieldValue == nil || fieldValue == "" {
			err := field.Set(db.Statement.ReflectValue, utils.String().Uuid(true))
			if err != nil {
				logger.Error("An error occurred while completing the entity's primary key.", err)
			}
		}
	}
}

func updateDb(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.SystemSettingEntity{},
		entity.UserEntity{})
	if err != nil {
		logger.Error("The system could not continue to run because it could not establish a connection with the database.", err)
		os.Exit(1)
	}
}
