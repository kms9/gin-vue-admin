import service from '@/utils/request'

// @Tags McCorp
// @Summary 创建McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.McCorp true "创建McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mcCorp/createMcCorp [post]
export const createMcCorp = (data) => {
     return service({
         url: "/mcCorp/createMcCorp",
         method: 'post',
         data
     })
 }


// @Tags McCorp
// @Summary 删除McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.McCorp true "删除McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mcCorp/deleteMcCorp [delete]
 export const deleteMcCorp = (data) => {
     return service({
         url: "/mcCorp/deleteMcCorp",
         method: 'delete',
         data
     })
 }

// @Tags McCorp
// @Summary 删除McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mcCorp/deleteMcCorp [delete]
 export const deleteMcCorpByIds = (data) => {
     return service({
         url: "/mcCorp/deleteMcCorpByIds",
         method: 'delete',
         data
     })
 }

// @Tags McCorp
// @Summary 更新McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.McCorp true "更新McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mcCorp/updateMcCorp [put]
 export const updateMcCorp = (data) => {
     return service({
         url: "/mcCorp/updateMcCorp",
         method: 'put',
         data
     })
 }


// @Tags McCorp
// @Summary 用id查询McCorp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.McCorp true "用id查询McCorp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mcCorp/findMcCorp [get]
 export const findMcCorp = (params) => {
     return service({
         url: "/mcCorp/findMcCorp",
         method: 'get',
         params
     })
 }


// @Tags McCorp
// @Summary 分页获取McCorp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取McCorp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mcCorp/getMcCorpList [get]
 export const getMcCorpList = (params) => {
     return service({
         url: "/mcCorp/getMcCorpList",
         method: 'get',
         params
     })
 }