package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baiduHotAPI = "https://top.baidu.com/api/board?platform=wise&tab=realtime"
)

// BaiduProvider 百度热搜数据提供者
type BaiduProvider struct{}

// NewBaiduProvider 创建新的百度热搜提供者
func NewBaiduProvider() *BaiduProvider {
	return &BaiduProvider{}
}

// baiduResponse 百度接口响应结构
type baiduResponse struct {
	Data struct {
		Cards []struct {
			Content []struct {
				Desc     string `json:"desc"`      // 描述
				HotScore string `json:"hotScore"`  // 热度值
				URL      string `json:"url"`       // 链接
				Query    string `json:"query"`     // 标题
				Index    int    `json:"index"`     // 排名
			} `json:"content"`
		} `json:"cards"`
	} `json:"data"`
}

// GetHotData 获取百度热搜数据
func (b *BaiduProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", baiduHotAPI, nil)
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
	var baiduResp baiduResponse
	if err := json.Unmarshal(body, &baiduResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	if len(baiduResp.Data.Cards) > 0 {
		for _, item := range baiduResp.Data.Cards[0].Content {
			result = append(result, HotData{
				Title:    item.Query,
				URL:      item.URL,
				Hot:      item.HotScore,
				Desc:     item.Desc,
				Index:    item.Index,
				Platform: b.GetPlatformName(),
			})
		}
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (b *BaiduProvider) GetPlatformName() string {
	return "百度"
}
