package tables

import "gorm.io/gorm"

// Pv 页面访问量
type Pv struct {
	gorm.Model
	Ip  string `json:"ip"`  // id地址
	Url string `json:"url"` // url
}
