// Code generated by goctl. DO NOT EDIT.
package types

type SearchRequset struct {
	OrderID string `form:"orderID"`
}

type SearchResponse struct {
	OrderID  string `json:"orderID"`
	Status   int    `json:"status"`
	Username string `json:"username"`
}
