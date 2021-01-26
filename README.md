# lemon-cloud-service
Lemon cloud is a private cloud storage system. This repo is the back-end API service code.

## 名词说明
plugin 插件分为两种。以下两种都可以称之为插件。但是业务和使用场景有区别。
1. ext 系统扩展：
   扩展lemon-cloud系统基础功能的程序。可以让系统为其增加设置项，
   也可以拥有仅WEB管理端可以使用的简单UI。例如：S3OSS存储扩展
2. app 应用程序： 
   基于云存储和访问会话的应用程序。应该带有跨平台的终端UI（暂定基于uni-app来实现）。
   通常APP是面对最终使用者的而非系统管理员。例如：.epub文件阅读器。