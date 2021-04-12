package request

import "gin-vue-admin/model"

type McCorpSearch struct{
    model.McCorp
    PageInfo
}