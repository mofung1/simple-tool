package autoload

// Wechat 微信配置
type Wechat struct {
	AppID     string `mapstructure:"appid" json:"appid" yaml:"appid"`
	AppSecret string `mapstructure:"secret" json:"secret" yaml:"secret"`
}
