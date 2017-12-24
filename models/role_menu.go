package models

type RoleMenu struct {
	RoleId int64 `orm:"column(role_id)"`
	Menuid int64 `orm:"column(menuid)"`
}
