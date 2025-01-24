package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	qqNewsHotAPI = "https://r.inews.qq.com/gw/event/hot_ranking_list"
)

// QQNewsProvider 腾讯新闻热榜数据提供者
type QQNewsProvider struct{}

// NewQQNewsProvider 创建新的腾讯新闻热榜提供者
func NewQQNewsProvider() *QQNewsProvider {
	return &QQNewsProvider{}
}

// qqNewsResponse 腾讯新闻接口响应结构
type qqNewsResponse struct {
	Code int `json:"ret"`
	Data struct {
		List []struct {
			Title    string `json:"title"`    // 标题
			URL      string `json:"url"`      // 链接
			HotValue int    `json:"hotValue"` // 热度值
			Summary  string `json:"summary"`  // 摘要
			Source   string `json:"source"`   // 来源
		} `json:"list"`
	} `json:"data"`
}

// GetHotData 获取腾讯新闻热榜数据
func (q *QQNewsProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", qqNewsHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://news.qq.com/")
	
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()
	
	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}
	
	// 解析响应数据
	var qqResp qqNewsResponse
	if err := json.Unmarshal(body, &qqResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if qqResp.Code != 0 {
		return nil, fmt.Errorf("腾讯新闻API返回错误码: %d", qqResp.Code)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range qqResp.Data.List {
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      fmt.Sprintf("%d热度", item.HotValue),
			Desc:     fmt.Sprintf("[%s] %s", item.Source, item.Summary),
			Index:    i + 1,
			Platform: q.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (q *QQNewsProvider) GetPlatformName() string {
	return "腾讯新闻"
}
