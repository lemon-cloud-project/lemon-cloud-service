package service

import (
	"os"
	"sync"
)

type SystemExtensionService struct {
	// 扩展程序的唯一Key和端口号之间的映射关系，如果扩展程序没有运行，那么存储端口号为0
	extensionWithPortMapping map[string]uint16
	// 扩展程序的唯一Key和进程实例之间的映射关系
	extensionWithProcessMapping map[string]*os.Process
}

var systemExtensionInstance *SystemExtensionService
var systemExtensionInitOnce sync.Once

func SystemExtension() *SystemExtensionService {
	systemExtensionInitOnce.Do(func() {
		systemExtensionInstance = &SystemExtensionService{}
	})
	return systemExtensionInstance
}

// 通过手动上传包来安装某个扩展程序，安装之后必须激活之后才可以使用
func (i *SystemExtensionService) InstallFromUpload() {

}

// 通过从应用商店下载来安装某个扩展程序，安装之后必须激活之后才可以使用
func (i *SystemExtensionService) InstallFromStore() {

}

// 无论是手动上传、还是从商店下载最后都通过该函数对得到的包进行安装
func (i *SystemExtensionService) InstallExtensionPackage(extensionPackagePath string) {

}

// 卸载某个扩展程序
func (i *SystemExtensionService) UnInstall(extensionKey string) {

}

// 激活某个扩展，激活的过程实际上就是将扩展程序的进程启动起来的过程
func (i *SystemExtensionService) Activate(extensionKey string) {

}

// 禁用某个扩展，禁用的过程实际上就是将某个扩展程序的进程杀死掉的过程
func (i *SystemExtensionService) Disable(extensionKey string) {

}

// 检查扩展程序列表，如果有缺失从集群主节点进行同步。暂未实现
func (i *SystemExtensionService) FixAndSynchronizeExtensionList() {

}
