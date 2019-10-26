// 作者 sukaifei
// 日期 2019/09/07

package model

type Article struct {
	Id int64 `gorm:"AUTO_INCREMENT;column:beast_id"`
	// 创建时间
	CreateTime int64 `gorm:"column:create_time"`
	// 内容
	HtmlText string `gorm:"column:html_text"`
	MdText   string `gorm:"column:md_text"`
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

type ArticleAdd struct {
	// 内容
	MdText string `json:"md_text" validate:"required"`
	// 内容
	HtmlText string `json:"html_text" validate:"required"`
	// 标题
	Title string `json:"title" validate:"required"`
	// 所属用户
	UserId int64 `json:"-"`
	// 所属栏目
	TagCode string `json:"tag" validate:"required"`
}

type ArticleInfo struct {
	Id int64 `json:"id" ddb:"id"`
	// 标题
	Title string `json:"title" ddb:"title"`
	// 正文
	MdText   string `json:"-" ddb:"md_text"`
	HtmlText string `json:"-" ddb:"html_text"`
	Text     string `json:"text"`
	// 所属用户昵称
	Author string `json:"author" ddb:"author"`
	// 所属用户头像
	Avatar string `json:"avatar" ddb:"avatar"`
	// 所属栏目
	TagName string `json:"tag" ddb:"tag"`
	// 阅读量
	PV int64 `json:"pv" ddb:"pv"`
	// 创建时间
	CreateTimeS int64  `json:"-" ddb:"create_time"`
	CreateTime  string `json:"create_time"`
}

type ArticleListItem struct {
	Id int64 `json:"id" ddb:"id"`
	// 标题
	Title string `json:"title" ddb:"title"`
	// 所属用户昵称
	Author string `json:"author" ddb:"author"`
	// 所属用户头像
	Avatar string `json:"avatar" ddb:"avatar"`
	// 所属栏目
	TagName string `json:"tag" ddb:"tag"`
	// 阅读量
	PV int64 `json:"pv" ddb:"pv"`
	// 创建时间
	CreateTimeS int64  `json:"-" ddb:"create_time"`
	CreateTime  string `json:"create_time"`
}

type ArticleList struct {
	Page   *Page              `json:"page"`
	Result []*ArticleListItem `json:"result"`
}

type MyArticleListReq struct {
	ParamPage
	Acc string `form:"acc"`
}

type ArticleListReq struct {
	ParamPage
	Tag string `form:"tag"`
}
