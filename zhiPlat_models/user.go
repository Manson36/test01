package zhiPlat_models

type SocialAccountType struct {
	Wechat string `json:"wechat"`
	QQ     string `json:"qq"`
	Weibo  string `json:"weibo"`
}

type User struct {
	ID         IDType            `json:"id"`
	Mobile     string            `json:"mobile" gorm:"pq_comment:用户手机号"`
	Email      string            `json:"email" gorm:"pq_comment:用户邮箱"`
	NickName   string            `json:"nickName" gorm:"pq_comment:用户昵称"`
	RealName   string            `json:"realName" gorm:"pq_comment:用户真实姓名"`
	Password   string            `json:"password" gorm:"pq_comment:加盐后的密码"`
	Salt       string            `json:"salt" gorm:"pq_comment:密码的盐"`
	Avatar     IDType            `json:"avatar" gorm:"pq_comment:用户头像关联的id"`
	AvatarPath string            `json:"avatarPath" gorm:"pq_comment:用户头像的路径"`
	Gender     int8              `json:"gender" gorm:"用户性别：0未知1男2女"`
	SocialAcc  SocialAccountType `json:"socialAcc" gorm:""`
}
