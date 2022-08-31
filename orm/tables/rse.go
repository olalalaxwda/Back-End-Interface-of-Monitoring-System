package tables

import "gorm.io/gorm"

// Rse 监听资源加载错误
type Rse struct {
	gorm.Model
	TargetClassName string `json:"targetClassName"` // 监听资源加载错误标签的类名
	Url             string `json:"url"`             // 资源错误url
}
