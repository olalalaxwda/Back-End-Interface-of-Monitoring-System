package tables

import "gorm.io/gorm"

// Ps 用户停留信息
type Ps struct {
	gorm.Model
	Page        string `json:"page"`        // 页名
	StayTime    string `json:"stayTime"`    // 页面存在时长
	VisibleTime string `json:"visibleTime"` // 页面显示时长
}
