package tables

import "gorm.io/gorm"

// Cpstate 异常发生请求成功失败时组件状态
type Cpstate struct {
	gorm.Model
	Name    string `json:"name"`    // 组件名称
	Data    string `json:"data"`    // 数据
	ErrorId uint   `json:"errorId"` // 异常id
	Type    string `json:"type"`    // 异常类型
}
