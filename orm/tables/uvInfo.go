package tables

import "gorm.io/gorm"

// Uv 独立访客数
type Uv struct {
	gorm.Model
	Ip  string `json:"ip"`  // ip地址
	Url string `json:"url"` // url
}
