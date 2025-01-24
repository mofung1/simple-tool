package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	neteaseHotAPI = "https://m.163.com/fe/api/hot/news/flow"
)

// NeteaseProvider 网易新闻热榜数据提供者
type NeteaseProvider struct{}

// NewNeteaseProvider 创建新的网易新闻热榜提供者
func NewNeteaseProvider() *NeteaseProvider {
	return &NeteaseProvider{}
}

// neteaseResponse 网易新闻响应结构体
type neteaseResponse struct {
	Code int    `json:"code"` // 响应码
	Msg  string `json:"msg"`  // 响应信息
	Data struct {
		List []struct {
			DocID      string `json:"docid"`      // 文档ID
			Title      string `json:"title"`      // 标题
			ImgSrc     string `json:"imgsrc"`     // 图片链接
			Source     string `json:"source"`     // 来源
			PTime      string `json:"ptime"`      // 发布时间
			CreateTime string `json:"createTime"` // 创建时间
			URL        string `json:"url"`        // 文章URL
		} `json:"list"`
	} `json:"data"`
}

// GetHotData 获取网易新闻热榜数据
func (n *NeteaseProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", neteaseHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://m.163.com/")

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
	var neteaseResp neteaseResponse
	if err := json.Unmarshal(body, &neteaseResp); err != nil {
		fmt.Printf("网易解析数据失败，错误: %v\n", err)
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}

	// 检查响应状态
	if neteaseResp.Code != 200 {
		return nil, fmt.Errorf("网易新闻API返回错误码: %d, 错误信息: %s", neteaseResp.Code, neteaseResp.Msg)
	}

	// 转换为通用格式
	var result HotDataList
	for idx, item := range neteaseResp.Data.List {
		result = append(result, HotData{
			Title:    item.Title,
			URL:      item.URL,
			Hot:      item.Source,
			Desc:     item.ImgSrc,
			Index:    idx + 1,
			Platform: n.GetPlatformName(),
		})
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("未找到任何热榜数据")
	}

	return result, nil
}

// GetPlatformName 获取平台名称
func (n *NeteaseProvider) GetPlatformName() string {
	return "网易新闻"
}
