package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	smzdmHotAPI = "https://api.smzdm.com/v1/list?category=all&type=good_price&offset=0&limit=20"
)

// SmzdmProvider 什么值得买热门数据提供者
type SmzdmProvider struct{}

// NewSmzdmProvider 创建新的什么值得买热门提供者
func NewSmzdmProvider() *SmzdmProvider {
	return &SmzdmProvider{}
}

// smzdmResponse 什么值得买接口响应结构
type smzdmResponse struct {
	Error_code string `json:"error_code"`
	Error_msg  string `json:"error_msg"`
	Data       struct {
		List []struct {
			ArticleID   string `json:"article_id"`   // 文章ID
			ArticleURL  string `json:"article_url"`  // 文章链接
			Title       string `json:"title"`        // 标题
			Price       string `json:"price"`        // 价格
			PriceString string `json:"price_string"` // 价格字符串
			Mall        string `json:"mall"`         // 商城
			ZhiCount    int    `json:"zhi_count"`   // 值评数
			BuzhiCount  int    `json:"buzhi_count"` // 不值评数
			CommentCount int   `json:"comment_count"`// 评论数
			Description string `json:"description"`  // 描述
		} `json:"list"`
	} `json:"data"`
}

// GetHotData 获取什么值得买热门数据
func (s *SmzdmProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", smzdmHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://www.smzdm.com/")
	
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
	var smzdmResp smzdmResponse
	if err := json.Unmarshal(body, &smzdmResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if smzdmResp.Error_code != "0" {
		return nil, fmt.Errorf("什么值得买API返回错误: %s", smzdmResp.Error_msg)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range smzdmResp.Data.List {
		// 构建描述信息
		desc := fmt.Sprintf("[%s] %s", item.Mall, item.Description)
		if len(desc) > 100 {
			desc = desc[:100] + "..."
		}
		
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.ArticleURL,
			Hot:      fmt.Sprintf("值%d 不值%d 评论%d", item.ZhiCount, item.BuzhiCount, item.CommentCount),
			Desc:     desc,
			Index:    i + 1,
			Platform: s.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (s *SmzdmProvider) GetPlatformName() string {
	return "什么值得买"
}
