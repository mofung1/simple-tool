package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	wereadHotAPI = "https://weread.qq.com/web/bookListInCategory/rising?maxIndex=0&rank=1"
)

// WereadProvider 微信读书热门数据提供者
type WereadProvider struct{}

// NewWereadProvider 创建新的微信读书热门提供者
func NewWereadProvider() *WereadProvider {
	return &WereadProvider{}
}

// wereadResponse 微信读书接口响应结构
type wereadResponse struct {
	Books []struct {
		BookInfo struct {
			BookID      string  `json:"bookId"`      // 图书ID
			Title       string  `json:"title"`       // 书名
			Author      string  `json:"author"`      // 作者
			Cover       string  `json:"cover"`       // 封面
			Intro       string  `json:"intro"`       // 简介
			Category    string  `json:"category"`    // 分类
			Price       float64 `json:"price"`       // 价格
			Rating      float64 `json:"rating"`      // 评分
			RatingCount int     `json:"ratingCount"` // 评分人数
		} `json:"bookInfo"`
		ReadCount     int `json:"readCount"`     // 阅读人数
		ReadingCount  int `json:"readingCount"`  // 正在阅读人数
		ReviewCount   int `json:"reviewCount"`   // 评论数
	} `json:"books"`
}

// GetHotData 获取微信读书热门数据
func (w *WereadProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", wereadHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://weread.qq.com/")
	
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
	var wereadResp wereadResponse
	if err := json.Unmarshal(body, &wereadResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range wereadResp.Books {
		// 构建描述信息
		desc := fmt.Sprintf("[%s] %s", item.BookInfo.Category, item.BookInfo.Intro)
		if len(desc) > 100 {
			desc = desc[:100] + "..."
		}
		
		result = append(result, HotData{
			Title:    fmt.Sprintf("%s - %s", item.BookInfo.Title, item.BookInfo.Author),
			URL:      fmt.Sprintf("https://weread.qq.com/web/reader/%s", item.BookInfo.BookID),
			Hot:      fmt.Sprintf("评分%.1f(%d人) 阅读%d人", item.BookInfo.Rating, item.BookInfo.RatingCount, item.ReadCount),
			Desc:     desc,
			Index:    i + 1,
			Platform: w.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (w *WereadProvider) GetPlatformName() string {
	return "微信读书"
}
