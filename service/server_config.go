package service

import (
	"fmt"
	"github.com/lemon-cloud-project/lemon-cloud-commons-golang/logger"
	"github.com/lemon-cloud-project/lemon-cloud-commons-golang/utils"
	"github.com/lemon-cloud-project/lemon-cloud-service/define"
	"github.com/lemon-cloud-project/lemon-cloud-service/model"
	"sync"
)

type ServerConfigService struct {
	ConfigFileLoadState bool
	ConfigObject        *model.ServerConfig
}

var serverConfigServiceInstance *ServerConfigService
var serverConfigServiceInitOnce sync.Once

func ServerConfig() *ServerConfigService {
	serverConfigServiceInitOnce.Do(func() {
		serverConfigServiceInstance = &ServerConfigService{}
		serverConfigServiceInstance.ConfigFileLoadState = false
		serverConfigServiceInstance.ConfigObject = &model.ServerConfig{}
	})
	return serverConfigServiceInstance
}

// Test whether MySQL can connect successfully
func (i *ServerConfigService) TestConnectMySQL(config *model.ServerConfig) error {
	// TODO
	return nil
}

// Test whether Redis can connect successfully
func (i *ServerConfigService) TestConnectRedis(config *model.ServerConfig) error {
	// TODO
	return nil
}

// Save configuration object
// If MySQL and redis test the connection successfully,
// the server configuration object will be saved to the JSON file on the local disk
func (i *ServerConfigService) Save(config *model.ServerConfig) error {
	if err := i.TestConnectMySQL(config); err == nil {
		return err
	}
	if err := i.TestConnectRedis(config); err == nil {
		return err
	}
	i.ConfigObject = config
	return i.SaveServerConfigToLocalJsonFile()
}

// Get the path of the configuration file
func (i *ServerConfigService) getConfigFilePath() string {
	return utils.IO().GetRuntimePath(define.ServerConfig().ServerConfigFileName())
}

// Get the server configuration object. If it does not exist, read it from the disk configuration file
func (i *ServerConfigService) GetServerConfig() *model.ServerConfig {
	if i.ConfigFileLoadState {
		return i.ConfigObject
	}
	if utils.IO().PathExists(i.getConfigFilePath()) {
		err := utils.IO().JsonFileToStruct(i.getConfigFilePath(), i.ConfigObject)
		if err != nil {
			logger.Error("Error reading configuration file locally.", err)
		} else {
			i.ConfigFileLoadState = true
		}
	} else {
		logger.Warn("No configuration file found, please complete the configuration and continue")
		err := i.SaveServerConfigToLocalJsonFile()
		if err != nil {
			logger.Error("Error generating empty configuration file", err)
		} else {
			logger.Warn("An empty configuration file has been generated in the directory where the program is located.")
			logger.Warn("You can modify the file manually or modify it in the visual UI")
		}
	}
	return i.ConfigObject
}

// Save the server configuration object as a local disk file
func (i ServerConfigService) SaveServerConfigToLocalJsonFile() error {
	return utils.IO().StructToJsonFile(i.getConfigFilePath(), i.ConfigObject)
}

// Get the connection URL of MySQL
func (i *ServerConfigService) GetMySQLConnectionUrl() string {
	mySQLConfig := i.GetServerConfig().MySQLConfig
	if mySQLConfig.Host == "" || mySQLConfig.Port == "" || mySQLConfig.Username == "" ||
		mySQLConfig.Password == "" || mySQLConfig.Schema == "" {
		return ""
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mySQLConfig.Username, mySQLConfig.Password,
		mySQLConfig.Host, mySQLConfig.Port, mySQLConfig.Schema)
}
