package datamodels

import "time"

type AdminUser struct {
	ID           int64      `json:"id, string"`
	NickName     string     `json:"nickName" gorm:"type:varchar(50);not null;default:'';pq_comment:用户昵称"`
	Email        string     `json:"email" grom:"type:varchar(200);not null;unique;pq_comment:用户邮箱"`
	Password     string     `json:"password" gorm:"type:varchar(200); not null;pq_comment:用户密码"`
	Salt         string     `json:"salt" gorm:"type:varchar(100);not null;pq_comment:加密密码的盐"`
	AdminType    int        `json:"adminType" gorm:"type:smallint;not null;default:2;pq_comment:管理用户的类型，1表示管理员，2表示普通用户"`
	PlatformType int        `json:"platformType" gorm:"type smallint;pq_comment:管理员所处的平台类型"`
	CreatedAt    *time.Time `json:"createAt" gorm:"type:timestamptz;not null;default:now();pq_comment:该用户的创建时间"`
	UpdatedAt    *time.Time `json:"updateAt" gorm:"type:timestamptz;default:now();pq_comment:该用户的更新时间"`
	RemovedAt    *time.Time `json:"removed" gorm:"type:timestamptz;pq_comment:该用户的移除时间"`
	Removed      bool       `json:"removed" gorm:"pq_comment:该用户是否被移除"`
	Disable      bool       `json:"disable" gorm:"pq_comment:用户是否被禁用"`
	DisableAt    *time.Time `json:"disableAt" gorm:"type:timestamptz;pq_comment:该用户被禁用的时间"`
}
