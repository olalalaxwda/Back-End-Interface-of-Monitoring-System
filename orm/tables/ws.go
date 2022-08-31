package tables

import "gorm.io/gorm"

// Ws 组件变化超时
type Ws struct {
	gorm.Model
	City string `json:"city"` // ip地址
}
