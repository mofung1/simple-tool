package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	toutiaoHotAPI = "https://www.toutiao.com/hot-event/hot-board/?origin=toutiao_pc"
)

// ToutiaoProvider 今日头条热榜数据提供者
type ToutiaoProvider struct{}

// NewToutiaoProvider 创建新的今日头条热榜提供者
func NewToutiaoProvider() *ToutiaoProvider {
	return &ToutiaoProvider{}
}

// toutiaoResponse 今日头条接口响应结构
type toutiaoResponse struct {
	Data []struct {
		Title    string `json:"Title"`    // 标题
		URL      string `json:"Url"`      // 链接
		HotValue int    `json:"HotValue"` // 热度值
		ClusterID string `json:"ClusterID"` // 聚合ID
	} `json:"data"`
}

// GetHotData 获取今日头条热榜数据
func (t *ToutiaoProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", toutiaoHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.toutiao.com/")
	
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
	var toutiaoResp toutiaoResponse
	if err := json.Unmarshal(body, &toutiaoResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range toutiaoResp.Data {
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      fmt.Sprintf("%d", item.HotValue),
			Desc:     "", // 头条热榜API没有提供描述信息
			Index:    i + 1,
			Platform: t.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (t *ToutiaoProvider) GetPlatformName() string {
	return "头条"
}
