package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	lolHotAPI = "https://bbs.nga.cn/app_api.php?__lib=post&__act=list&__output=14&fid=152&page=1&lite=js&noprefix"
)

// LOLProvider 英雄联盟热门数据提供者
type LOLProvider struct{}

// NewLOLProvider 创建新的英雄联盟热门提供者
func NewLOLProvider() *LOLProvider {
	return &LOLProvider{}
}

// lolResponse 英雄联盟接口响应结构
type lolResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Result struct {
		Data []struct {
			Tid       int    `json:"tid"`        // 帖子ID
			Subject   string `json:"subject"`     // 标题
			Author    string `json:"author"`      // 作者
			Postdate  int64  `json:"postdate"`   // 发帖时间
			Replies   int    `json:"replies"`    // 回复数
			Views     int    `json:"views"`      // 浏览数
			Content   string `json:"content"`    // 内容
			LastPost  int64  `json:"lastpost"`   // 最后回复时间
			TopLevel  int    `json:"top_level"`  // 置顶等级
			Type      int    `json:"type"`       // 帖子类型
		} `json:"data"`
	} `json:"result"`
}

// GetHotData 获取英雄联盟热门数据
func (l *LOLProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", lolHotAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://bbs.nga.cn/")
	
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
	var lolResp lolResponse
	if err := json.Unmarshal(body, &lolResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 检查响应状态
	if lolResp.Code != 0 {
		return nil, fmt.Errorf("英雄联盟API返回错误: %s", lolResp.Msg)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range lolResp.Result.Data {
		// 构建描述信息
		desc := fmt.Sprintf("by %s", item.Author)
		if len(item.Content) > 0 {
			if len(item.Content) > 100 {
				desc += "\n" + item.Content[:100] + "..."
			} else {
				desc += "\n" + item.Content
			}
		}
		
		// 添加帖子类型标签
		typeStr := ""
		switch item.Type {
		case 1:
			typeStr = "[公告] "
		case 2:
			typeStr = "[精华] "
		}
		
		result = append(result, HotData{
			Title:    typeStr + item.Subject,
			URL:      fmt.Sprintf("https://bbs.nga.cn/read.php?tid=%d", item.Tid),
			Hot:      fmt.Sprintf("回复%d 浏览%d", item.Replies, item.Views),
			Desc:     desc,
			Index:    i + 1,
			Platform: l.GetPlatformName(),
		})
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (l *LOLProvider) GetPlatformName() string {
	return "英雄联盟"
}
