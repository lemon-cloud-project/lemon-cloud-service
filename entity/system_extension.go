package entity

import (
	"github.com/lemon-cloud-project/lemon-cloud-service/base"
)

type SystemExtensionEntity struct {
	base.Entity
	Key                 string `gorm:"unique;type:string;size:64;comment:'The unique readable key of the system extension. Read from the extension package. If it already exists in the database, installation is not allowed.';" json:"key"`
	Name                string `gorm:"type:string;size:128;comment:'The name of the extension, read from the extension package';" json:"name"`
	Author              string `gorm:"type:string;size:256;comment:'Author of the extension, read from the extension package';" json:"author"`
	Introduce           string `gorm:"type:string;size:10240;comment:'The introduction of the extension, read from the extension package'" json:"introduce"`
	Version             string `gorm:"type:string;size:64;comment:'The version number of the currently installed program, read from the extension package'" json:"version"`
	ExpectedActiveState bool   `gorm:"type:bool;comment:'Expected extended activation state, after opening the program will keep the process alive as much as possible.'" json:"expected_active_state"`
	InstallationMethod  string `gorm:"type:string;size:16;comment:'How to install this extension. Enum type. UPLOAD: Locally uploaded packages;STORE: Download and install the package from the store;'" json:"installation_method"`
	// 源代码类型，平台根据此字段决定通过什么命令来启动扩展程序进程。为了以后插件的多语言开发而做准备的。
	// 目前支持的类型枚举及示意：GO: go语言标准格式
	SourceCodeType string `gorm:"type:string;size:16;comment:'Source code type. The platform will decide what command to use to start a new process based on the source code type. GO：Go Language Standard';" json:"source_code_type"`
}
