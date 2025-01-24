package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	douyinHotAPI = "https://www.douyin.com/aweme/v1/web/hot/search/list/"
)

// DouyinProvider 抖音热点数据提供者
type DouyinProvider struct{}

// NewDouyinProvider 创建新的抖音热点提供者
func NewDouyinProvider() *DouyinProvider {
	return &DouyinProvider{}
}

// douyinResponse 抖音接口响应结构
type douyinResponse struct {
	Data struct {
		WordList []struct {
			Word     string `json:"word"`      // 标题
			HotValue int    `json:"hot_value"` // 热度值
			Link     string `json:"link"`      // 链接
		} `json:"word_list"`
	} `json:"data"`
}

// GetHotData 获取抖音热点数据
func (d *DouyinProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", douyinHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.douyin.com/")
	
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
	var douyinResp douyinResponse
	if err := json.Unmarshal(body, &douyinResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range douyinResp.Data.WordList {
		result = append(result, HotData{
			Title:    item.Word,
			URL:      item.Link,
			Hot:      fmt.Sprintf("%d", item.HotValue),
			Desc:     "", // 抖音热搜API没有提供描述信息
			Index:    i + 1,
			Platform: d.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (d *DouyinProvider) GetPlatformName() string {
	return "抖音"
}
