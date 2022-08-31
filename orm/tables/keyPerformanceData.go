package tables

import "gorm.io/gorm"

// KeyPerformanceData 性能指标
type KeyPerformanceData struct {
	gorm.Model
	DnsTime           string `json:"dnsTime"`           // dns时间
	DomReady          string `json:"domReady"`          // domReady时间
	FirstContentPaint string `json:"firstContentPaint"` // 记录页面第一次绘制像素的时间
	FirstPaintTime    string `json:"firstPaintTime"`    // 首次绘制
	WhiteScreenTime   string `json:"whiteScreenTime"`   // 白屏时间
}
