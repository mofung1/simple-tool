package hot

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	ithomeHotAPI = "https://m.ithome.com/rankm/"
)

// IthomeProvider IT之家热榜数据提供者
type IthomeProvider struct{}

// NewIthomeProvider 创建新的IT之家热榜提供者
func NewIthomeProvider() *IthomeProvider {
	return &IthomeProvider{}
}

// replaceLink 处理IT之家的链接
func replaceLink(url string) string {
	re := regexp.MustCompile(`[html|live]/(\d+)\.htm`)
	matches := re.FindStringSubmatch(url)
	if len(matches) > 1 {
		id := matches[1]
		if len(id) > 3 {
			return fmt.Sprintf("https://www.ithome.com/0/%s/%s.htm", id[:3], id[3:])
		}
	}
	return url
}

// GetHotData 获取IT之家热榜数据
func (i *IthomeProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", ithomeHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Referer", "https://m.ithome.com/")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %v", err)
	}

	// 转换为通用格式
	var result HotDataList
	doc.Find(".rank-box .placeholder").Each(func(idx int, s *goquery.Selection) {
		// 获取链接
		href, exists := s.Find("a").Attr("href")
		if !exists {
			return
		}

		// 获取标题
		title := strings.TrimSpace(s.Find(".plc-title").Text())

		// 获取评论数
		commentText := strings.TrimSpace(s.Find(".review-num").Text())
		commentCount := 0
		if re := regexp.MustCompile(`\d+`); re.MatchString(commentText) {
			commentCount, _ = strconv.Atoi(re.FindString(commentText))
		}

		// 获取封面图
		cover, _ := s.Find("img").Attr("data-original")

		// 处理链接
		url := replaceLink(href)

		// 添加到结果中
		if title != "" && url != "" {
			result = append(result, HotData{
				Title:    title,
				URL:      url,
				Hot:      fmt.Sprintf("%d评论", commentCount),
				Desc:     fmt.Sprintf("%s", cover),
				Index:    idx + 1,
				Platform: i.GetPlatformName(),
			})
		}
	})

	if len(result) == 0 {
		return nil, fmt.Errorf("未找到任何热榜数据")
	}

	return result, nil
}

// GetPlatformName 获取平台名称
func (i *IthomeProvider) GetPlatformName() string {
	return "IT之家"
}
