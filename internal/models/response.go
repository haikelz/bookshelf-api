package models

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type APISuccessResponse struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"`
}

type APIErrorResponse struct {
	BaseResponse
}
