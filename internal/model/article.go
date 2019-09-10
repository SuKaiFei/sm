// 作者 sukaifei
// 日期 2019/09/07

package model

type Article struct {
	Id int64 `gorm:"AUTO_INCREMENT;column:beast_id"`
	// 创建时间
	CreateTime int64 `gorm:"column:create_time"`
	// 内容
	Content string `gorm:"column:content"`
	// 阅读量
	Pv int64 `gorm:"column:pv"`
	// 点赞数
	ThumbsUp int64 `gorm:"column:thumbs_up"`
	// 标题
	Title string `gorm:"column:title"`
	// 所属用户
	UserId int64 `gorm:"column:user_id"`
	// 所属栏目
	TagCode string `gorm:"column:tag_code"`
}
