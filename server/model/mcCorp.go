// 自动生成模板McCorp
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type McCorp struct {
      global.GVA_MODEL
      ContactSecret  string `json:"contactSecret" form:"contactSecret" gorm:"column:contact_secret;comment:企业外部联系人secret;type:char;"`
      EmployeeSecret  string `json:"employeeSecret" form:"employeeSecret" gorm:"column:employee_secret;comment:企业通讯录secret;type:char;"`
      EncodingAesKey  string `json:"encodingAesKey" form:"encodingAesKey" gorm:"column:encoding_aes_key;comment:回调消息加密串;type:char;"`
      EventCallback  string `json:"eventCallback" form:"eventCallback" gorm:"column:event_callback;comment:事件回调地址;type:varchar(255);size:255;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:企业名称;type:varchar(255);size:255;"`
      SocialCode  string `json:"socialCode" form:"socialCode" gorm:"column:social_code;comment:企业代码(企业统一社会信用代码);type:char;"`
      TenantId  int `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;type:int;size:10;"`
      Token  string `json:"token" form:"token" gorm:"column:token;comment:回调token;type:char;"`
      WxCorpid  string `json:"wxCorpid" form:"wxCorpid" gorm:"column:wx_corpid;comment:企业微信ID;type:char;"`
}


func (McCorp) TableName() string {
  return "mc_corp"
}

