package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	csdnHotAPI = "https://blog.csdn.net/phoenix/web/blog/hot-rank"
)

// CSDNProvider CSDN热榜数据提供者
type CSDNProvider struct{}

// NewCSDNProvider 创建新的CSDN热榜提供者
func NewCSDNProvider() *CSDNProvider {
	return &CSDNProvider{}
}

// csdnResponse CSDN接口响应结构
type csdnResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			ArticleTitle       string `json:"articleTitle"`       // 文章标题
			ArticleDetailURL   string `json:"articleDetailUrl"`   // 文章链接
			ViewCount         int    `json:"viewCount"`          // 浏览数
			CommentCount      int    `json:"commentCount"`       // 评论数
			DiggCount         int    `json:"diggCount"`          // 点赞数
			NickName          string `json:"nickName"`           // 作者昵称
			ArticleDesc       string `json:"articleDesc"`        // 文章描述
			PicList          []string `json:"picList"`          // 图片列表
		} `json:"list"`
	} `json:"data"`
}

// GetHotData 获取CSDN热榜数据
func (c *CSDNProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", csdnHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://blog.csdn.net/")
	
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
	var csdnResp csdnResponse
	if err := json.Unmarshal(body, &csdnResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if csdnResp.Code != 200 {
		return nil, fmt.Errorf("CSDN API返回错误: %s", csdnResp.Message)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range csdnResp.Data.List {
		result = append(result, HotData{
			Title:    item.ArticleTitle,
			URL:      item.ArticleDetailURL,
			Hot:      fmt.Sprintf("浏览%d 评论%d 点赞%d", item.ViewCount, item.CommentCount, item.DiggCount),
			Desc:     fmt.Sprintf("[%s] %s", item.NickName, item.ArticleDesc),
			Index:    i + 1,
			Platform: c.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (c *CSDNProvider) GetPlatformName() string {
	return "CSDN"
}
