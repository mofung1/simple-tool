package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	weiboHotSearchAPI = "https://weibo.com/ajax/side/hotSearch"
)

// WeiboProvider 微博热搜数据提供者
type WeiboProvider struct{}

// NewWeiboProvider 创建新的微博热搜提供者
func NewWeiboProvider() *WeiboProvider {
	return &WeiboProvider{}
}

// weiboResponse 微博接口响应结构
type weiboResponse struct {
	Data struct {
		RealTime []struct {
			Note    string `json:"note"`     // 标题
			URL     string `json:"url"`      // 链接
			Hot     int    `json:"hot"`      // 热度
			Rank    int    `json:"rank"`     // 排名
			Content string `json:"content"`   // 内容描述
		} `json:"realtime"`
	} `json:"data"`
	Ok int `json:"ok"`
}

// GetHotData 获取微博热搜数据
func (w *WeiboProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", weiboHotSearchAPI, nil)
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
	var weiboResp weiboResponse
	if err := json.Unmarshal(body, &weiboResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for _, item := range weiboResp.Data.RealTime {
		result = append(result, HotData{
			Title:    item.Note,
			URL:      "https://s.weibo.com" + item.URL,
			Hot:      fmt.Sprintf("%d", item.Hot),
			Desc:     item.Content,
			Index:    item.Rank,
			Platform: w.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (w *WeiboProvider) GetPlatformName() string {
	return "微博"
}
