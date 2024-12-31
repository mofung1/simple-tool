package autoload

type LogConfig struct {
	Level      string `mapstructure:"level"`
	LogDir     string `mapstructure:"log_dir"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
