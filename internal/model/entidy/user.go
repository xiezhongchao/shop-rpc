package entidy

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mobile        string `gorm:"type:varchar(11);unique;not null;comment:手机号"`
	Password      string `gorm:"type:varchar(30);unique;not null;comment:密码"`
	Email         string `gorm:"type:varchar(30);unique;not null;comment:邮箱"`
	NickName      string `gorm:"type:varchar(30);unique;not null;comment:昵称"`
	Gender        int    `gorm:"type:int(2);unique;not null;comment:性别"`
	AvatarUrl     string `gorm:"type:varchar(30);unique;not null;comment:头像地址"`
	AccountStatus int    `gorm:"type:int(2);unique;not null;comment:账号状态"`
}
