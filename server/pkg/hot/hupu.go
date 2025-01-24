package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	hupuHotAPI = "https://bbs.hupu.com/api/v1/hot-news"
)

// HupuProvider 虎扑热门数据提供者
type HupuProvider struct{}

// NewHupuProvider 创建新的虎扑热门提供者
func NewHupuProvider() *HupuProvider {
	return &HupuProvider{}
}

// hupuResponse 虎扑接口响应结构
type hupuResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			ID            int    `json:"id"`             // 帖子ID
			Title         string `json:"title"`          // 标题
			URL          string `json:"url"`            // 链接
			ReplyCount   int    `json:"reply_count"`    // 回复数
			LightCount   int    `json:"light_count"`    // 亮了数
			CreateTime   string `json:"create_time"`    // 创建时间
			ForumName    string `json:"forum_name"`     // 版块名称
			UserNickname string `json:"user_nickname"`  // 发帖人昵称
		} `json:"list"`
	} `json:"data"`
}

// GetHotData 获取虎扑热门数据
func (h *HupuProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", hupuHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://bbs.hupu.com/")
	
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
	var hupuResp hupuResponse
	if err := json.Unmarshal(body, &hupuResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if hupuResp.Code != 0 {
		return nil, fmt.Errorf("虎扑API返回错误: %s", hupuResp.Message)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range hupuResp.Data.List {
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      fmt.Sprintf("回复%d 亮了%d", item.ReplyCount, item.LightCount),
			Desc:     fmt.Sprintf("[%s] by %s", item.ForumName, item.UserNickname),
			Index:    i + 1,
			Platform: h.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (h *HupuProvider) GetPlatformName() string {
	return "虎扑"
}
