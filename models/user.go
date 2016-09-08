package models

type User struct {
	Id            int64
	Account       string
	Password      string
	NickName      string
	CreateTime    int64
	ModifyTime    int64
	LastLoginTime int64
	Permissions   int64
}
