package usecase

import (
	"fmt"
	"gin-vue-admin/test/lib"
	"gin-vue-admin/test/wxlib"
	"github.com/xen0n/go-workwx"
)
func GetUsers()  {
	info, err:=wxlib.WorkwxClient.ListAllDepts()
	if err!=nil {
		lib.LLog.Error(info)
	}

	userList:= make([]*workwx.UserInfo, 0)
	useMapList:=make(map[string]*workwx.UserInfo, 0)
	for _, v:=range info{
		fmt.Printf("%+v", v)
		deptUsers, err := wxlib.WorkwxClient.ListUsersByDeptID(v.ID, true)
		if err!=nil {
			lib.LLog.Error(info)
			continue
		}
		userList = append(userList, deptUsers...)
	}

	for _, v:=range userList {
		useMapList[v.Email] = v

	}

	for _, v:=range useMapList {
		fmt.Println("\n")
		fmt.Printf("%+v", v)
	}
}

func DoWork()  {
	GetUsers()
}