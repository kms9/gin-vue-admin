package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMcCorpRouter(Router *gin.RouterGroup) {
	McCorpRouter := Router.Group("mcCorp").Use(middleware.OperationRecord())
	{
		McCorpRouter.POST("createMcCorp", v1.CreateMcCorp)   // 新建McCorp
		McCorpRouter.DELETE("deleteMcCorp", v1.DeleteMcCorp) // 删除McCorp
		McCorpRouter.DELETE("deleteMcCorpByIds", v1.DeleteMcCorpByIds) // 批量删除McCorp
		McCorpRouter.PUT("updateMcCorp", v1.UpdateMcCorp)    // 更新McCorp
		McCorpRouter.GET("findMcCorp", v1.FindMcCorp)        // 根据ID获取McCorp
		McCorpRouter.GET("getMcCorpList", v1.GetMcCorpList)  // 获取McCorp列表
	}
}
