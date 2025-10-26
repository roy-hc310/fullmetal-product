package dto

type Query struct {
	Page   string `form:"page" json:"page"`
	Size   string `form:"size" json:"size"`
	Cursor string `form:"cursor" json:"cursor"`

	Keyword string `form:"keyword" json:"keyword"`
	SortBy  string `form:"sort" json:"sort_by"`
}

type Pagination struct {
	Page       string `json:"page"`
	Size       string `json:"size"`
	TotalItems string `json:"total_items"`
	NextCursor string `json:"next_cursor"`
}

type ResponseObject struct {
	Data      interface{} `json:"data"`
	TraceID   string      `json:"trace_id"`
	Succeeded bool        `json:"succeeded"`
	Errors    []string    `json:"errors"`
}

type ResponseArray struct {
	ResponseObject
	Pagination
}
