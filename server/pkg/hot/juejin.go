package hot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	juejinHotAPI = "https://api.juejin.cn/recommend_api/v1/article/recommend_all_feed"
)

// JuejinProvider 掘金热榜数据提供者
type JuejinProvider struct{}

// NewJuejinProvider 创建新的掘金热榜提供者
func NewJuejinProvider() *JuejinProvider {
	return &JuejinProvider{}
}

// juejinResponse 掘金接口响应结构
type juejinResponse struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Data   []struct {
		Item struct {
			ArticleInfo struct {
				Title        string `json:"title"`         // 标题
				BriefContent string `json:"brief_content"` // 简介
			} `json:"article_info"`
			ArticleID    string `json:"article_id"`    // 文章ID
			ViewCount    int    `json:"view_count"`    // 浏览数
			DiggCount    int    `json:"digg_count"`    // 点赞数
			CommentCount int    `json:"comment_count"` // 评论数
		} `json:"item_info"`
	} `json:"data"`
}

// GetHotData 获取掘金热榜数据
func (j *JuejinProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求体
	requestBody := map[string]interface{}{
		"id_type":       2,
		"sort_type":     3, // 按热度排序
		"cate_id":       "6809637767543259144", // 全部分类
		"cursor":        "0",
		"limit":         20,
	}
	
	// 将请求体转换为JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("创建请求体失败: %v", err)
	}
	
	// 创建请求
	req, err := http.NewRequest("POST", juejinHotAPI, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://juejin.cn/")
	
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
	var juejinResp juejinResponse
	if err := json.Unmarshal(body, &juejinResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if juejinResp.ErrNo != 0 {
		return nil, fmt.Errorf("掘金API返回错误: %s", juejinResp.ErrMsg)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range juejinResp.Data {
		result = append(result, HotData{
			Title:    item.Item.ArticleInfo.Title,
			URL:      fmt.Sprintf("https://juejin.cn/post/%s", item.Item.ArticleID),
			Hot:      fmt.Sprintf("点赞%d 评论%d 浏览%d", item.Item.DiggCount, item.Item.CommentCount, item.Item.ViewCount),
			Desc:     item.Item.ArticleInfo.BriefContent,
			Index:    i + 1,
			Platform: j.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (j *JuejinProvider) GetPlatformName() string {
	return "掘金"
}
