// Package hot 提供热点数据获取功能
package hot

// HotData 定义热点数据的基本结构
type HotData struct {
	Title    string `json:"title"`     // 标题
	URL      string `json:"url"`       // 链接
	Hot      string `json:"hot"`       // 热度值
	Desc     string `json:"desc"`      // 描述
	Index    int    `json:"index"`     // 排序索引
	Platform string `json:"platform"`   // 来源平台
}

// HotDataList 热点数据列表
type HotDataList []HotData

// HotDataProvider 定义热点数据提供者接口
type HotDataProvider interface {
	// GetHotData 获取热点数据
	GetHotData() (HotDataList, error)
	// GetPlatformName 获取平台名称
	GetPlatformName() string
}
