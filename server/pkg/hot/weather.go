package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	weatherAlarmAPI = "https://api.wolfx.jp/weather_alarm.json"
)

// WeatherProvider 天气预警数据提供者
type WeatherProvider struct{}

// NewWeatherProvider 创建新的天气预警提供者
func NewWeatherProvider() *WeatherProvider {
	return &WeatherProvider{}
}

// weatherResponse 天气预警接口响应结构
type weatherResponse struct {
	UpdateTime int64 `json:"update_time"`
	Data      []struct {
		ID          string    `json:"id"`           // 预警ID
		Title       string    `json:"title"`        // 标题
		Type        string    `json:"type"`         // 预警类型
		Level       string    `json:"level"`        // 预警等级
		Status      string    `json:"status"`       // 状态
		IssueTime   string    `json:"issue_time"`   // 发布时间
		Location    string    `json:"location"`     // 预警地区
		Description string    `json:"description"`  // 详细描述
		Source      string    `json:"source"`       // 发布源
	} `json:"data"`
}

// GetHotData 获取天气预警数据
func (w *WeatherProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", weatherAlarmAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	
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
	var weatherResp weatherResponse
	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range weatherResp.Data {
		// 构建标题
		title := fmt.Sprintf("%s %s", item.Location, item.Title)
		
		// 构建描述信息
		desc := fmt.Sprintf("[%s发布] %s", item.Source, item.Description)
		if len(desc) > 100 {
			desc = desc[:100] + "..."
		}
		
		// 构建热度信息（预警等级和类型）
		hot := fmt.Sprintf("%s级 %s", item.Level, item.Type)
		
		result = append(result, HotData{
			Title:    title,
			URL:      "", // 天气预警暂时没有详情页面
			Hot:      hot,
			Desc:     desc,
			Index:    i + 1,
			Platform: w.GetPlatformName(),
		})
		
		// 只返回最近20条记录
		if i >= 19 {
			break
		}
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (w *WeatherProvider) GetPlatformName() string {
	return "天气预警"
}
