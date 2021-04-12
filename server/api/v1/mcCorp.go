package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags McCorp
// @Summary 创建McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.McCorp true "创建McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mcCorp/createMcCorp [post]
func CreateMcCorp(c *gin.Context) {
	var mcCorp model.McCorp
	_ = c.ShouldBindJSON(&mcCorp)
	if err := service.CreateMcCorp(mcCorp); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags McCorp
// @Summary 删除McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.McCorp true "删除McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mcCorp/deleteMcCorp [delete]
func DeleteMcCorp(c *gin.Context) {
	var mcCorp model.McCorp
	_ = c.ShouldBindJSON(&mcCorp)
	if err := service.DeleteMcCorp(mcCorp); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags McCorp
// @Summary 批量删除McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mcCorp/deleteMcCorpByIds [delete]
func DeleteMcCorpByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteMcCorpByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags McCorp
// @Summary 更新McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.McCorp true "更新McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mcCorp/updateMcCorp [put]
func UpdateMcCorp(c *gin.Context) {
	var mcCorp model.McCorp
	_ = c.ShouldBindJSON(&mcCorp)
	if err := service.UpdateMcCorp(mcCorp); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags McCorp
// @Summary 用id查询McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.McCorp true "用id查询McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mcCorp/findMcCorp [get]
func FindMcCorp(c *gin.Context) {
	var mcCorp model.McCorp
	_ = c.ShouldBindQuery(&mcCorp)
	if err, remcCorp := service.GetMcCorp(mcCorp.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remcCorp": remcCorp}, c)
	}
}

// @Tags McCorp
// @Summary 分页获取McCorp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.McCorpSearch true "分页获取McCorp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mcCorp/getMcCorpList [get]
func GetMcCorpList(c *gin.Context) {
	var pageInfo request.McCorpSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMcCorpInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
