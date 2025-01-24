package hot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	earthquakeAPI = "https://api.wolfx.jp/cenc_eqlist.json"
)

// EarthquakeProvider 地震速报数据提供者
type EarthquakeProvider struct{}

// NewEarthquakeProvider 创建新的地震速报提供者
func NewEarthquakeProvider() *EarthquakeProvider {
	return &EarthquakeProvider{}
}

// earthquakeResponse 地震速报接口响应结构
type earthquakeResponse struct {
	UpdateTime int64 `json:"update_time"`
	Data      []struct {
		ID           string  `json:"id"`            // 地震ID
		Time         string  `json:"time"`          // 发生时间
		Magnitude    float64 `json:"magnitude"`     // 震级
		Depth        float64 `json:"depth"`         // 震源深度
		Location     string  `json:"location"`      // 震中位置
		Latitude     float64 `json:"latitude"`      // 纬度
		Longitude    float64 `json:"longitude"`     // 经度
		Level        string  `json:"level"`         // 烈度
		ReportStatus int     `json:"report_status"` // 报告状态
	} `json:"data"`
}

// GetHotData 获取地震速报数据
func (e *EarthquakeProvider) GetHotData() (HotDataList, error) {
	// 创建HTTP客户端
	client := &http.Client{}
	
	// 创建请求
	req, err := http.NewRequest("GET", earthquakeAPI, nil)
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
	var eqResp earthquakeResponse
	if err := json.Unmarshal(body, &eqResp); err != nil {
		return nil, fmt.Errorf("解析数据失败: %v", err)
	}
	
	// 转换为通用格式
	var result HotDataList
	for i, item := range eqResp.Data {
		// 构建标题
		title := fmt.Sprintf("%s %s %.1f级地震", item.Time, item.Location, item.Magnitude)
		
		// 构建描述信息
		desc := fmt.Sprintf("震源深度: %.1f千米\n位置坐标: %.4f, %.4f", 
			item.Depth, item.Latitude, item.Longitude)
		
		// 构建地震地图URL
		mapURL := fmt.Sprintf("https://www.google.com/maps?q=%.4f,%.4f", 
			item.Latitude, item.Longitude)
		
		// 构建热度信息
		hot := fmt.Sprintf("%.1f级", item.Magnitude)
		if item.Level != "" {
			hot += fmt.Sprintf(" %s", item.Level)
		}
		
		result = append(result, HotData{
			Title:    title,
			URL:      mapURL,
			Hot:      hot,
			Desc:     desc,
			Index:    i + 1,
			Platform: e.GetPlatformName(),
		})
		
		// 只返回最近20条记录
		if i >= 19 {
			break
		}
	}
	
	return result, nil
}

// GetPlatformName 获取平台名称
func (e *EarthquakeProvider) GetPlatformName() string {
	return "地震速报"
}
