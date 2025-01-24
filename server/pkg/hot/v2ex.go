package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	v2exHotAPI = "https://www.v2ex.com/api/v2/topics/hot.json"
)

// V2exProvider V2EX热门数据提供者
type V2exProvider struct{}

// NewV2exProvider 创建新的V2EX热门提供者
func NewV2exProvider() *V2exProvider {
	return &V2exProvider{}
}

// v2exResponse V2EX接口响应结构
type v2exResponse struct {
	Success bool `json:"success"`
	Data    []struct {
		ID            int    `json:"id"`              // 话题ID
		Title         string `json:"title"`           // 标题
		Content       string `json:"content"`         // 内容
		URL          string `json:"url"`             // 链接
		Replies      int    `json:"replies"`         // 回复数
		LastReplyID  int    `json:"last_reply_id"`   // 最后回复ID
		LastTouched  int    `json:"last_touched"`    // 最后更新时间
		CreatedAt    int    `json:"created"`         // 创建时间
		Node         struct {
			Name     string `json:"name"`     // 节点名称
			Title    string `json:"title"`    // 节点标题
		} `json:"node"`
		Member       struct {
			Username string `json:"username"` // 用户名
		} `json:"member"`
	} `json:"data"`
}

// GetHotData 获取V2EX热门数据
func (v *V2exProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", v2exHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://www.v2ex.com/")
	
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
	var v2exResp v2exResponse
	if err := json.Unmarshal(body, &v2exResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if !v2exResp.Success {
		return nil, fmt.Errorf("V2EX API返回失败状态")
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range v2exResp.Data {
		desc := fmt.Sprintf("[%s] %s", item.Node.Title, item.Content)
		if len(desc) > 100 {
			desc = desc[:100] + "..."
		}
		
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      fmt.Sprintf("%d回复", item.Replies),
			Desc:     desc,
			Index:    i + 1,
			Platform: v.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (v *V2exProvider) GetPlatformName() string {
	return "V2EX"
}
