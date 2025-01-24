package hot

import "fmt"

// Platform 定义支持的平台类型
type Platform string

const (
	// 主流资讯平台
	PlatformWeibo    Platform = "weibo"    // 微博热搜
	PlatformZhihu    Platform = "zhihu"    // 知乎热榜
	PlatformBaidu    Platform = "baidu"    // 百度热搜
	PlatformDouyin   Platform = "douyin"   // 抖音热点
	PlatformToutiao  Platform = "toutiao"  // 今日头条
	PlatformBilibili Platform = "bilibili" // 哔哩哔哩
	PlatformIthome   Platform = "ithome"   // IT之家
	
	// 新闻平台
	PlatformNetease  Platform = "netease"  // 网易新闻
	PlatformQQNews   Platform = "qqnews"   // 腾讯新闻
	PlatformSina     Platform = "sina"     // 新浪新闻
	PlatformThePaper Platform = "thepaper" // 澎湃新闻
	
	// 技术社区
	PlatformJuejin   Platform = "juejin"   // 掘金
	PlatformV2ex     Platform = "v2ex"     // V2EX
	PlatformCSDN     Platform = "csdn"     // CSDN
	PlatformGithub   Platform = "github"   // GitHub
	
	// 生活娱乐
	PlatformDouban   Platform = "douban"   // 豆瓣电影
	PlatformSmzdm    Platform = "smzdm"    // 什么值得买
	PlatformHupu     Platform = "hupu"     // 虎扑
	PlatformWeread   Platform = "weread"   // 微信读书
	
	// 游戏相关
	PlatformMiyoushe Platform = "miyoushe" // 米游社
	PlatformGenshin  Platform = "genshin"  // 原神
	PlatformLOL      Platform = "lol"      // 英雄联盟
	PlatformHonkai   Platform = "honkai"   // 崩坏：星穹铁道
	
	// 特色服务
	PlatformEarthquake Platform = "earthquake" // 地震速报
	PlatformWeather    Platform = "weather"    // 天气预警
)

// NewHotDataProvider 创建热点数据提供者的工厂方法
func NewHotDataProvider(platform Platform) (HotDataProvider, error) {
	switch platform {
	// 主流资讯平台
	case PlatformWeibo:
		return NewWeiboProvider(), nil
	case PlatformZhihu:
		return NewZhihuProvider(), nil
	case PlatformBaidu:
		return NewBaiduProvider(), nil
	case PlatformDouyin:
		return NewDouyinProvider(), nil
	case PlatformToutiao:
		return NewToutiaoProvider(), nil
	case PlatformBilibili:
		return NewBilibiliProvider(), nil
	case PlatformIthome:
		return NewIthomeProvider(), nil
		
	// 新闻平台
	case PlatformNetease:
		return NewNeteaseProvider(), nil
	case PlatformQQNews:
		return NewQQNewsProvider(), nil
	case PlatformSina:
		return NewSinaNewsProvider(), nil
	case PlatformThePaper:
		return NewThePaperProvider(), nil
		
	// 技术社区
	case PlatformJuejin:
		return NewJuejinProvider(), nil
	case PlatformV2ex:
		return NewV2exProvider(), nil
	case PlatformCSDN:
		return NewCSDNProvider(), nil
	case PlatformGithub:
		return NewGithubProvider(), nil
		
	// 生活娱乐
	case PlatformDouban:
		return NewDoubanProvider(), nil
	case PlatformSmzdm:
		return NewSmzdmProvider(), nil
	case PlatformHupu:
		return NewHupuProvider(), nil
	case PlatformWeread:
		return NewWereadProvider(), nil
		
	// 游戏相关
	case PlatformMiyoushe:
		return NewMiyousheProvider(), nil
	case PlatformGenshin:
		return NewGenshinProvider(), nil
	case PlatformLOL:
		return NewLOLProvider(), nil
	case PlatformHonkai:
		return NewHonkaiProvider(), nil
		
	// 特色服务
	case PlatformEarthquake:
		return NewEarthquakeProvider(), nil
	case PlatformWeather:
		return NewWeatherProvider(), nil
		
	default:
		return nil, fmt.Errorf("不支持的平台类型: %s", platform)
	}
}

// GetAllPlatforms 获取所有支持的平台
func GetAllPlatforms() []Platform {
	return []Platform{
		// 主流资讯平台
		PlatformWeibo,
		PlatformZhihu,
		PlatformBaidu,
		PlatformDouyin,
		PlatformToutiao,
		PlatformBilibili,
		PlatformIthome,
		
		// 新闻平台
		PlatformNetease,
		PlatformQQNews,
		PlatformSina,
		PlatformThePaper,
		
		// 技术社区
		PlatformJuejin,
		PlatformV2ex,
		PlatformCSDN,
		PlatformGithub,
		
		// 生活娱乐
		PlatformDouban,
		PlatformSmzdm,
		PlatformHupu,
		PlatformWeread,
		
		// 游戏相关
		PlatformMiyoushe,
		PlatformGenshin,
		PlatformLOL,
		PlatformHonkai,
		
		// 特色服务
		PlatformEarthquake,
		PlatformWeather,
	}
}

// GetAllHotData 获取所有平台的热点数据
func GetAllHotData() (map[Platform]HotDataList, error) {
	platforms := GetAllPlatforms()
	result := make(map[Platform]HotDataList)
	
	for _, platform := range platforms {
		provider, err := NewHotDataProvider(platform)
		if err != nil {
			return nil, fmt.Errorf("创建%s平台提供者失败: %v", platform, err)
		}
		
		data, err := provider.GetHotData()
		if err != nil {
			// 这里我们选择记录错误但继续执行，而不是直接返回错误
			// 这样即使某个平台失败，其他平台的数据仍然可以返回
			result[platform] = HotDataList{}
			continue
		}
		
		result[platform] = data
	}
	
	return result, nil
}
