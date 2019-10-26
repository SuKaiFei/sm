package model

const (
	// 默认时间格式
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

type TopStatus int32

const (
	// 正文格式
	MdText   = "markdown"
	HtmlText = "html"

	// 置顶状态
	_ TopStatus = iota
	NotTop
	IsTop
)

// Page pagination
type Page struct {
	Num   int `json:"num"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

// Page Param
type ParamPage struct {
	Pn int32 `json:"pn" form:"pn" default:"1"`
	Ps int32 `json:"ps" form:"ps" default:"10"`
}

type CommonInfoReq struct {
	Id int64 `json:"id" form:"id"`
}
