// Code generated by goctl. DO NOT EDIT.
package types

type SignupRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender,option=male|femal|secret,default=secret"`
}

type SignupResponse struct {
	Message string `json:"message"`
}