package model

type MySQLConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
}

type RedisConfig struct {
	Network  string `json:"network"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

type ServerConfig struct {
	MySQLConfig MySQLConfig `json:"mysql_config"`
	RedisConfig RedisConfig `json:"redis_config"`
	// 本节点的工作区路径，工作区用来放一些系统运行所需的数据及文件，如：系统扩展包、应用程序包等
	WorkspacePath string `json:"workspace_path"`
}
