package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	githubTrendingAPI = "https://api.github.com/search/repositories"
)

// GithubProvider GitHub Trending数据提供者
type GithubProvider struct{}

// NewGithubProvider 创建新的GitHub Trending提供者
func NewGithubProvider() *GithubProvider {
	return &GithubProvider{}
}

// githubResponse GitHub接口响应结构
type githubResponse struct {
	Items []struct {
		Name        string `json:"name"`         // 仓库名称
		FullName    string `json:"full_name"`    // 完整仓库名
		HTMLURL     string `json:"html_url"`     // 仓库链接
		Description string `json:"description"`   // 仓库描述
		Language    string `json:"language"`      // 主要语言
		StarCount   int    `json:"stargazers_count"` // star数
		ForkCount   int    `json:"forks_count"`      // fork数
		Owner       struct {
			Login string `json:"login"`          // 作者名称
		} `json:"owner"`
	} `json:"items"`
}

// GetHotData 获取GitHub Trending数据
func (g *GithubProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", githubTrendingAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置查询参数：获取过去一周内创建的，按star数排序的仓库
	q := req.URL.Query()
	q.Add("q", "created:>"+getLastWeekDate())
	q.Add("sort", "stars")
	q.Add("order", "desc")
	q.Add("per_page", "20")
	req.URL.RawQuery = q.Encode()
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	
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
	var githubResp githubResponse
	if err := json.Unmarshal(body, &githubResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range githubResp.Items {
		lang := item.Language
		if lang == "" {
			lang = "未知语言"
		}
		
		result = append(result, HotData{
			Title:    fmt.Sprintf("%s/%s", item.Owner.Login, item.Name),
			URL:      item.HTMLURL,
			Hot:      fmt.Sprintf("Star:%d Fork:%d", item.StarCount, item.ForkCount),
			Desc:     fmt.Sprintf("[%s] %s", lang, item.Description),
			Index:    i + 1,
			Platform: g.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (g *GithubProvider) GetPlatformName() string {
	return "GitHub"
}

// getLastWeekDate 获取一周前的日期，格式为YYYY-MM-DD
func getLastWeekDate() string {
	now := time.Now()
	lastWeek := now.AddDate(0, 0, -7)
	return lastWeek.Format("2006-01-02")
}
