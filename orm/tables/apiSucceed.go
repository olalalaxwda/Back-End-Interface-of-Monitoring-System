package tables

import "gorm.io/gorm"

// ApiSucceed api请求成功
type ApiSucceed struct {
	gorm.Model
	Status string `json:"status"` // 状态码
	Url    string `json:"url"`    // 路径
}
