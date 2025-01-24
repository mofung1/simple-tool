package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	sinaNewsHotAPI = "https://interface.sina.cn/wap_api/hot_list.d.json"
)

// SinaNewsProvider 新浪新闻热榜数据提供者
type SinaNewsProvider struct{}

// NewSinaNewsProvider 创建新的新浪新闻热榜提供者
func NewSinaNewsProvider() *SinaNewsProvider {
	return &SinaNewsProvider{}
}

// sinaNewsResponse 新浪新闻接口响应结构
type sinaNewsResponse struct {
	Status struct {
		Code int `json:"code"`
	} `json:"status"`
	Data struct {
		List []struct {
			Title    string `json:"title"`    // 标题
			URL      string `json:"url"`      // 链接
			ReadNum  int    `json:"read_num"` // 阅读数
			Media    string `json:"media"`    // 来源
			Abstract string `json:"abstract"` // 摘要
		} `json:"list"`
	} `json:"data"`
}

// GetHotData 获取新浪新闻热榜数据
func (s *SinaNewsProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", sinaNewsHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://news.sina.com.cn/")
	
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
	var sinaResp sinaNewsResponse
	if err := json.Unmarshal(body, &sinaResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if sinaResp.Status.Code != 0 {
		return nil, fmt.Errorf("新浪新闻API返回错误码: %d", sinaResp.Status.Code)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range sinaResp.Data.List {
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      fmt.Sprintf("%d阅读", item.ReadNum),
			Desc:     fmt.Sprintf("[%s] %s", item.Media, item.Abstract),
			Index:    i + 1,
			Platform: s.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (s *SinaNewsProvider) GetPlatformName() string {
	return "新浪新闻"
}
