package model

//Pager ...
type Pager struct {
	PageNo     int    `json:"pageNo"`
	PageSize   int    `json:"pageSize"`
	OrderField string `json:"orderField"`
	OrderASC   string `json:"orderASC"`
}
