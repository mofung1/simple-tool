package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	neteaseHotAPI = "https://m.163.com/fe/api/hot/news/flow"
)

// NeteaseProvider 网易新闻热榜数据提供者
type NeteaseProvider struct{}

// NewNeteaseProvider 创建新的网易新闻热榜提供者
func NewNeteaseProvider() *NeteaseProvider {
	return &NeteaseProvider{}
}

// neteaseResponse 网易新闻接口响应结构
type neteaseResponse struct {
	Code int `json:"code"`
	Data struct {
		Items []struct {
			Title    string `json:"title"`    // 标题
			URL      string `json:"url"`      // 链接
			HotValue int    `json:"hotValue"` // 热度值
			Source   string `json:"source"`   // 来源
			Digest   string `json:"digest"`   // 摘要
		} `json:"items"`
	} `json:"data"`
}

// GetHotData 获取网易新闻热榜数据
func (n *NeteaseProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", neteaseHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.163.com/")
	
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
	var neteaseResp neteaseResponse
	if err := json.Unmarshal(body, &neteaseResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if neteaseResp.Code != 0 {
		return nil, fmt.Errorf("网易新闻API返回错误码: %d", neteaseResp.Code)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range neteaseResp.Data.Items {
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      fmt.Sprintf("%d热度", item.HotValue),
			Desc:     fmt.Sprintf("[%s] %s", item.Source, item.Digest),
			Index:    i + 1,
			Platform: n.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (n *NeteaseProvider) GetPlatformName() string {
	return "网易新闻"
}
