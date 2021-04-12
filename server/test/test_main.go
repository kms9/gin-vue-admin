package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"gin-vue-admin/test/lib"
	"gin-vue-admin/test/usecase"
)

func main()  {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Redis()

	lib.LLog.Info("log Start")

	usecase.DoWork()
}

