// 作者 sukaifei
// 日期 2019/09/07

package model

type Tag struct {
	Id int64 `gorm:"AUTO_INCREMENT;column:beast_id"`
	// 创建时间
	CreateTime int64 `gorm:"column:create_time"`
	// 排序
	Sort int32 `gorm:"column:sort"`
	// 名称
	Name string `gorm:"column:name" ddb:"name"`
	// 代码
	Code string `gorm:"column:code" ddb:"code"`
}

type Tags struct {
	// 排序
	Sort int32 `gorm:"column:sort"  ddb:"sort" json:"sort,omitempty"`
	// 名称
	Name string `gorm:"column:name" ddb:"name" json:"name,omitempty"`
	// 代码
	Code string `gorm:"column:code" ddb:"code" json:"code,omitempty"`
}
