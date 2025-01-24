package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	zhihuHotAPI = "https://api.zhihu.com/topstory/hot-list"
)

// ZhihuProvider 知乎热榜数据提供者
type ZhihuProvider struct{}

// NewZhihuProvider 创建新的知乎热榜提供者
func NewZhihuProvider() *ZhihuProvider {
	return &ZhihuProvider{}
}

// zhihuResponse 知乎接口响应结构
type zhihuResponse struct {
	Data []struct {
		Target struct {
			Title      string `json:"title"`       // 标题
			URL        string `json:"url"`         // 链接
			ExcerptArea struct {
				Text string `json:"text"`          // 描述
			} `json:"excerpt_area"`
			MetricsArea struct {
				Text string `json:"text"`          // 热度
			} `json:"metrics_area"`
		} `json:"target"`
	} `json:"data"`
}

// GetHotData 获取知乎热榜数据
func (z *ZhihuProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", zhihuHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	
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
	var zhihuResp zhihuResponse
	if err := json.Unmarshal(body, &zhihuResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range zhihuResp.Data {
		result = append(result, HotData{
			Title:    item.Target.Title,
			URL:      item.Target.URL,
			Hot:      item.Target.MetricsArea.Text,
			Desc:     item.Target.ExcerptArea.Text,
			Index:    i + 1,
			Platform: z.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (z *ZhihuProvider) GetPlatformName() string {
	return "知乎"
}
