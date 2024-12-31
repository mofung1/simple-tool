package autoload

type JwtConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireTime string `mapstructure:"expire_time"`
}
