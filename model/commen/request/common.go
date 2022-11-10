package request

type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
	Keyword  int `json:"keyword" form:"keyword"`
}

type RequestID struct {
	ID uint `json:"id" form:"id"`
}
