package tables

import "gorm.io/gorm"

// ApiFailed api请求失败
type ApiFailed struct {
	gorm.Model
	Code    string `json:"code"`    // 请求异常状态
	Message string `json:"message"` // 请求异常信息
	Url     string `json:"url"`     // 请求异常url
}
