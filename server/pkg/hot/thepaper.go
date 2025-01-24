package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	thePaperHotAPI = "https://api.thepaper.cn/contentapi/nodeCont/getByNodeIdAndPage"
)

// ThePaperProvider 澎湃新闻热榜数据提供者
type ThePaperProvider struct{}

// NewThePaperProvider 创建新的澎湃新闻热榜提供者
func NewThePaperProvider() *ThePaperProvider {
	return &ThePaperProvider{}
}

// thePaperResponse 澎湃新闻接口响应结构
type thePaperResponse struct {
	Data struct {
		List []struct {
			Name      string `json:"name"`      // 标题
			URL       string `json:"url"`       // 链接
			ReadCount int    `json:"readCount"` // 阅读数
			Source    string `json:"source"`    // 来源
			Abstract  string `json:"abstract"`  // 摘要
		} `json:"list"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// GetHotData 获取澎湃新闻热榜数据
func (t *ThePaperProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", thePaperHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求参数
	q := req.URL.Query()
	q.Add("nodeId", "25949")  // 热门新闻节点ID
	q.Add("pageSize", "20")   // 每页数量
	q.Add("pageNum", "1")     // 页码
	req.URL.RawQuery = q.Encode()
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://www.thepaper.cn/")
	
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
	var paperResp thePaperResponse
	if err := json.Unmarshal(body, &paperResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if paperResp.Code != 0 {
		return nil, fmt.Errorf("澎湃新闻API返回错误: %s", paperResp.Message)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range paperResp.Data.List {
		result = append(result, HotData{
			Title:    item.Name,
			URL:      item.URL,
			Hot:      fmt.Sprintf("%d阅读", item.ReadCount),
			Desc:     fmt.Sprintf("[%s] %s", item.Source, item.Abstract),
			Index:    i + 1,
			Platform: t.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (t *ThePaperProvider) GetPlatformName() string {
	return "澎湃新闻"
}
