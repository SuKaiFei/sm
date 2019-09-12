// 作者 sukaifei
// 日期 2019/09/07

package model

type User struct {
	Id int64 `gorm:"AUTO_INCREMENT;column:id"`
	// 头像
	Avatar string `gorm:"column:avatar"`
	// 热度
	Heat string `gorm:"column:heat"`
	// 最后一次登录时间
	LastLoginTime int64 `gorm:"column:last_login_time"`
	// 昵称
	NickName string `gorm:"column:nickname"`
	// 第三方登录：OpenId
	OpenId string `gorm:"column:open_id"`
	// 角色
	Role string `gorm:"column:role"`
	// 性别
	Sex *string `gorm:"column:sex"`
	// 标志
	Sign string `gorm:"column:sign"`
	// 创建时间
	CreateTime int64 `gorm:"column:create_time"`
	// 主站点
	MainSite string `gorm:"column:main_site"`
	// 邮箱
	Email string `gorm:"column:email"`
	// 第三方登录：GitHub
	GithubId string `gorm:"column:github_id"`
	// 密码（BCrypt）
	Password string `gorm:"column:password"`
	// 会员级别
	Vip string `gorm:"column:vip"`
	// 验证中邮箱
	ValidateEmail int32 `gorm:"column:validate_email"`
}

type Login struct {
	UserName string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Registered struct {
	NickName string `json:"nickname" form:"nickname" validate:"required"`
	Sex      string `json:"sex" form:"sex" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserInfo struct {
	Id int64 `json:"id"`
	// 头像
	Avatar string `json:"avatar"`
	// 热度
	Heat string `json:"heat"`
	// 最后一次登录时间
	LastLoginTime int64 `json:"last_login_time"`
	// 昵称
	NickName string `json:"nickname"`
	// 角色
	Role string `json:"role"`
	// 性别
	Sex *string `json:"sex"`
	// 邮箱
	Email string `json:"email"`
	// 会员级别
	Vip string `json:"vip"`
}
