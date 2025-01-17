package autoload

type AppConfig struct {
	App           `mapstructure:"app"`
	*LogConfig    `mapstructure:"log"`
	*DBConfig     `mapstructure:"db"`
	*RedisConfig  `mapstructure:"redis"`
	*JwtConfig    `mapstructure:"jwt"`
	*Wechat       `mapstructure:"wechat"`
	*OpenAIConfig `mapstructure:"openai"`
}
