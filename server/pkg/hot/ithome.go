package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ithomeHotAPI = "https://api.ithome.com/json/newslist/news"
)

// IthomeProvider IT之家热榜数据提供者
type IthomeProvider struct{}

// NewIthomeProvider 创建新的IT之家热榜提供者
func NewIthomeProvider() *IthomeProvider {
	return &IthomeProvider{}
}

// ithomeResponse IT之家接口响应结构
type ithomeResponse struct {
	Newslist []struct {
		Title       string `json:"title"`       // 标题
		URL         string `json:"url"`         // 链接
		NewsID      string `json:"newsid"`      // 新闻ID
		Commentnum  string `json:"commentnum"`  // 评论数
		Description string `json:"description"` // 描述
	} `json:"newslist"`
}

// GetHotData 获取IT之家热榜数据
func (i *IthomeProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", ithomeHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.ithome.com/")
	
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
	var ithomeResp ithomeResponse
	if err := json.Unmarshal(body, &ithomeResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for idx, item := range ithomeResp.Newslist {
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      fmt.Sprintf("%s评论", item.Commentnum),
			Desc:     item.Description,
			Index:    idx + 1,
			Platform: i.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (i *IthomeProvider) GetPlatformName() string {
	return "IT之家"
}
