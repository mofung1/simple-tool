package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	doubanMovieAPI = "https://movie.douban.com/j/search_subjects?type=movie&tag=热门&sort=recommend&page_limit=20&page_start=0"
)

// DoubanProvider 豆瓣电影热门数据提供者
type DoubanProvider struct{}

// NewDoubanProvider 创建新的豆瓣电影热门提供者
func NewDoubanProvider() *DoubanProvider {
	return &DoubanProvider{}
}

// doubanResponse 豆瓣电影接口响应结构
type doubanResponse struct {
	Subjects []struct {
		Title    string  `json:"title"`     // 电影名称
		URL      string  `json:"url"`       // 链接
		Rate     string  `json:"rate"`      // 评分
		Cover    string  `json:"cover"`     // 封面图
		ID       string  `json:"id"`        // 电影ID
		IsNew    bool    `json:"is_new"`    // 是否新上映
		Playable bool    `json:"playable"`  // 是否可播放
		Cover_x  int     `json:"cover_x"`   // 封面宽度
		Cover_y  int     `json:"cover_y"`   // 封面高度
	} `json:"subjects"`
}

// GetHotData 获取豆瓣电影热门数据
func (d *DoubanProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", doubanMovieAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://movie.douban.com/")
	
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
	var doubanResp doubanResponse
	if err := json.Unmarshal(body, &doubanResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range doubanResp.Subjects {
		// 构建描述信息
		desc := fmt.Sprintf("评分: %s", item.Rate)
		if item.IsNew {
			desc += " [新上映]"
		}
		if item.Playable {
			desc += " [可播放]"
		}
		
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      item.Rate + "分",
			Desc:     desc,
			Index:    i + 1,
			Platform: d.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (d *DoubanProvider) GetPlatformName() string {
	return "豆瓣电影"
}
