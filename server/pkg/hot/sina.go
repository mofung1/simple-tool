package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	sinaNewsHotAPI = "https://newsapp.sina.cn/api/hotlist?newsId=HB-1-snhs/top_news_list-all"
)

// SinaNewsProvider 新浪新闻热搜数据提供者
type SinaNewsProvider struct{}

// NewSinaNewsProvider 创建新的新浪新闻热搜提供者
func NewSinaNewsProvider() *SinaNewsProvider {
	return &SinaNewsProvider{}
}

// sinaNewsResponse 新浪新闻热榜响应结构
type sinaNewsResponse struct {
	Status int    `json:"status"` // 状态码
	Msg    string `json:"msg"`    // 状态信息
	Data   struct {
		Data struct {
			HotList []struct {
				Title     string `json:"title"`      // 标题
				URL       string `json:"url"`        // 链接
				HotValue  string `json:"hotValue"`   // 热度值
				UniqueKey string `json:"unique_key"` // 唯一标识
			} `json:"hotList"`
		} `json:"data"`
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
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://sinanews.sina.cn/")
	
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

	// 打印原始响应以便调试
	fmt.Printf("新浪新闻原始响应: %s\n", string(body))
	
	// 解析响应数据
	var sinaResp sinaNewsResponse
	if err := json.Unmarshal(body, &sinaResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v, 响应内容: %s", err, string(body))
	}

	// 检查响应状态
	if sinaResp.Status != 0 || sinaResp.Msg != "success" {
		return nil, fmt.Errorf("新浪新闻API返回错误状态: %d, %s", sinaResp.Status, sinaResp.Msg)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range sinaResp.Data.Data.HotList {
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      item.HotValue,
			Index:    i + 1,
			Platform: s.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (s *SinaNewsProvider) GetPlatformName() string {
	return string(PlatformSina)
}
