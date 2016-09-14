package sqlite

import (
	"fmt"
	"fvCloud/models"
	"github.com/astaxie/beego/orm"
	"time"
)

//@Title 添加账号
func AddUser(account string, password string, nickname string, permissions int64) error {
	o := orm.NewOrm()
	tm := time.Now().UnixNano()
	user := &models.User{Account: account, Password: password,
		NickName: nickname, CreateTime: tm, ModifyTime: tm,
		LastLoginTime: tm, Permissions: permissions}
	qs := o.QueryTable("user")
	err := qs.Filter("account", account).One(user)
	if err == nil {
		return fmt.Errorf("existence account") //不能添加同账号的
	}
	_, err = o.Insert(user)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
