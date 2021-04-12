package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMcCorp
//@description: 创建McCorp记录
//@param: mcCorp model.McCorp
//@return: err error

func CreateMcCorp(mcCorp model.McCorp) (err error) {
	err = global.GVA_DB.Create(&mcCorp).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMcCorp
//@description: 删除McCorp记录
//@param: mcCorp model.McCorp
//@return: err error

func DeleteMcCorp(mcCorp model.McCorp) (err error) {
	err = global.GVA_DB.Delete(&mcCorp).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMcCorpByIds
//@description: 批量删除McCorp记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMcCorpByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.McCorp{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMcCorp
//@description: 更新McCorp记录
//@param: mcCorp *model.McCorp
//@return: err error

func UpdateMcCorp(mcCorp model.McCorp) (err error) {
	err = global.GVA_DB.Save(&mcCorp).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMcCorp
//@description: 根据id获取McCorp记录
//@param: id uint
//@return: err error, mcCorp model.McCorp

func GetMcCorp(id uint) (err error, mcCorp model.McCorp) {
	err = global.GVA_DB.Where("id = ?", id).First(&mcCorp).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMcCorpInfoList
//@description: 分页获取McCorp记录
//@param: info request.McCorpSearch
//@return: err error, list interface{}, total int64

func GetMcCorpInfoList(info request.McCorpSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.McCorp{})
    var mcCorps []model.McCorp
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&mcCorps).Error
	return err, mcCorps, total
}