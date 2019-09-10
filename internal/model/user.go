// 作者 sukaifei
// 日期 2019/09/07

package model

type User struct {
	Id int64 `gorm:"AUTO_INCREMENT;column:beast_id"`
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
	Sex string `gorm:"column:sex"`
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
	ValidateEmail string `gorm:"column:validate_email"`
}
