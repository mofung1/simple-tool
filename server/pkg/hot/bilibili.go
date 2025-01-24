package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	bilibiliHotAPI = "https://api.bilibili.com/x/web-interface/ranking/v2"
)

// BilibiliProvider B站热门数据提供者
type BilibiliProvider struct{}

// NewBilibiliProvider 创建新的B站热门提供者
func NewBilibiliProvider() *BilibiliProvider {
	return &BilibiliProvider{}
}

// bilibiliResponse B站接口响应结构
type bilibiliResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Title       string `json:"title"`         // 标题
			URI        string `json:"uri"`          // 链接
			Desc       string `json:"desc"`         // 描述
			Stat       struct {
				View     int `json:"view"`      // 播放量
				Like     int `json:"like"`      // 点赞数
				Reply    int `json:"reply"`     // 评论数
			} `json:"stat"`
			ShortLinkV2 string `json:"short_link_v2"` // 短链接
		} `json:"list"`
	} `json:"data"`
}

// GetHotData 获取B站热门数据
func (b *BilibiliProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", bilibiliHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.bilibili.com/")
	
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
	var biliResp bilibiliResponse
	if err := json.Unmarshal(body, &biliResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if biliResp.Code != 0 {
		return nil, fmt.Errorf("B站API返回错误: %s", biliResp.Message)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range biliResp.Data.List {
		// 优先使用短链接
		url := item.ShortLinkV2
		if url == "" {
			url = item.URI
		}
		
		result = append(result, HotData{
			Title:    item.Title,
			URL:      url,
			Hot:      fmt.Sprintf("%d播放 %d点赞", item.Stat.View, item.Stat.Like),
			Desc:     item.Desc,
			Index:    i + 1,
			Platform: b.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (b *BilibiliProvider) GetPlatformName() string {
	return "哔哩哔哩"
}
