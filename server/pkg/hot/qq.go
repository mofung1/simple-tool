package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

const (
	qqNewsHotAPI = "https://r.inews.qq.com/gw/event/hot_ranking_list?page_size=50"
)

// QQNewsProvider 腾讯新闻热榜数据提供者
type QQNewsProvider struct{}

// NewQQNewsProvider 创建新的腾讯新闻热榜提供者
func NewQQNewsProvider() *QQNewsProvider {
	return &QQNewsProvider{}
}

// qqNewsResponse 腾讯新闻接口响应结构
type qqNewsResponse struct {
	Ret    int    `json:"ret"`     // 响应码
	IDList []struct {
		IDsHash  string `json:"ids_hash"` // ID hash
		HasMore  int    `json:"has_more"` // 是否有更多
		NewsList []struct {
			ID        string `json:"id"`        // 新闻ID
			Title     string `json:"title"`     // 标题
			Abstract  string `json:"abstract"`  // 摘要
			Source    string `json:"source"`    // 来源
			Time      string `json:"time"`      // 时间
			Timestamp int64  `json:"timestamp"` // 时间戳
			HotEvent  struct {
				ID       string `json:"id"`       // 热点ID
				Ranking  int    `json:"ranking"`  // 排名
				Title    string `json:"title"`    // 热点标题
				HotScore int    `json:"hotScore"` // 热度值
			} `json:"hotEvent"`
		} `json:"newslist"`
	} `json:"idlist"`
}

// GetHotData 获取腾讯新闻热榜数据
func (q *QQNewsProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", qqNewsHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://news.qq.com/")

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
	var qqResp qqNewsResponse
	if err := json.Unmarshal(body, &qqResp); err != nil {
		fmt.Printf("腾讯新闻解析数据失败，错误: %v\n", err)
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}

	// 检查响应状态
	if qqResp.Ret != 0 {
		return nil, fmt.Errorf("腾讯新闻API返回错误码: %d", qqResp.Ret)
	}

	// 检查数据是否为空
	if len(qqResp.IDList) == 0 || len(qqResp.IDList[0].NewsList) == 0 {
		return nil, fmt.Errorf("未找到任何热榜数据")
	}

	// 转换为通用格式
	var result HotDataList
	for _, news := range qqResp.IDList[0].NewsList {
		if news.HotEvent.HotScore > 0 {
			hot := HotData{
				Title:    news.Title,
				URL:      fmt.Sprintf("https://view.inews.qq.com/a/%s", news.ID),
				Hot:      fmt.Sprintf("%d", news.HotEvent.HotScore),
				Desc:     news.Abstract,
				Index:    news.HotEvent.Ranking,
				Platform: q.GetPlatformName(),
			}
			result = append(result, hot)
		}
	}

	// 按照热度排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].Index < result[j].Index
	})

	return result, nil
}

// GetPlatformName 获取平台名称
func (q *QQNewsProvider) GetPlatformName() string {
	return "腾讯新闻"
}
