package models

import (
	"time"
)

type SysUser struct {
	UerID              int       `json:"user_id" gorm:"column:user_id;type:int;primary_key;auto_increment;comment:用户ID"`
	UserName           string    `json:"username" gorm:"column:username;type:varchar(255);default:null;comment:英文名"`
	UserNameCn         string    `json:"username_cn" gorm:"column:username_cn;type:varchar(255);default:null;comment:中文名"`
	NickName           string    `json:"nickname" gorm:"column:nickname;type:varchar(255);default:null;comment:昵称"`
	Password           string    `json:"password" gorm:"column:password;type:varchar(255);default:null;comment:密码"`
	Salt               string    `json:"salt" gorm:"column:salt;type:varchar(255);default:null;comment:随机盐"`
	Phone              string    `json:"phone" gorm:"column:phone;type:varchar(255);default:null;comment:手机"`
	Email              string    `json:"email" gorm:"column:email;type:varchar(255);default:null;comment:邮箱"`
	CreateTime         time.Time `json:"create_time" gorm:"column:create_time;type:datetime;default:current_timestamp;comment:创建时间"`
	UpdateTime         time.Time `json:"update_time" gorm:"column:update_time;type:datetime;default:null on update current_timestamp;comment:修改时间"`
	UpdatePasswordTime time.Time `json:"update_password_time" gorm:"column:update_password_time;type:datetime;default:null;comment:更新密码时间"`
	LockFlag           int       `json:"lock_flag" gorm:"column:lock_flag;default:0;comment:是否锁定 0:正常 9:锁定 2:冻结"`
}

type SysRole struct {
	RoleID     int       `json:"role_id" gorm:"column:role_id;type:int;primary_key;auto_increment;comment:角色ID"`
	RoleName   string    `json:"role_name" gorm:"column:role_name;type:varchar(255);default:null;comment:角色名字"`
	RoleCode   string    `json:"role_code" gorm:"column:role_code;type:varchar(255);default:null;comment:角色代码"`
	RoleDesc   string    `json:"role_desc" gorm:"column:role_desc;type:varchar(255);default:null;comment:角色说明"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:datetime;default:current_timestamp;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;type:datetime;default:null on update current_timestamp;comment:修改时间"`
}

type SysMenu struct {
	MenuID     int       `json:"menu_id" gorm:"column:menu_id;type:int;primary_key;auto_increment;comment:菜单ID"`
	Name       string    `json:"name" gorm:"column:name;type:varchar(32);default:null;comment:菜单名字"`
	Permission string    `json:"permission" gorm:"column:permission;type:varchar(32);default:null;comment:权限"`
	Path       string    `json:"path" gorm:"column:path;type:varchar(128);default:null;comment:路径"`
	ParentID   int       `json:"parent_id" gorm:"column:parent_id;type:int;default:null;comment:父菜单ID"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:datetime;default:current_timestamp;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;type:datetime;default:null on update current_timestamp;comment:修改时间"`
}

type SysUserRole struct {
	UserID int `json:"user_id" gorm:"column:user_id;type:int;primary_key;comment:用户ID"`
	RoleID int `json:"role_id" gorm:"column:role_id;type:int;primary_key;comment:角色ID"`
}

type SysRoleMenu struct {
	RoleID int `json:"role_id" gorm:"column:role_id;type:int;primary_key;comment:角色ID"`
	MenuID int `json:"menu_id" gorm:"column:menu_id;type:int;primary_key;comment:菜单ID"`
}
