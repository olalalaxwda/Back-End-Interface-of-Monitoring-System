package tables

import "gorm.io/gorm"

// Reject 监听未捕获的reject错误
type Reject struct {
	gorm.Model
	Reason string `json:"reason"` // 为捕获的reject内容
}
