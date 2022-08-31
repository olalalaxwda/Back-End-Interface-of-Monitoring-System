package tables

import "gorm.io/gorm"

// Jspe js同步异步异常
type Jspe struct {
	gorm.Model
	Filename string `json:"filename"` // 异常文件路径
	Lineno   string `json:"lineno"`   // 异常行号
	Message  string `json:"message"`  // 异常信息
}
