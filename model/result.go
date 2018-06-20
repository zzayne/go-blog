package model

//Result ...
type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

//PageResult ...
type PageResult struct {
	Result
	TotalCount int `json:"totalCount"`
}
