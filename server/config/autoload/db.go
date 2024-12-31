package autoload

type DBConfig struct {
	Type                string `mapstructure:"type"`
	Host                string `mapstructure:"host"`
	User                string `mapstructure:"user"`
	Password            string `mapstructure:"password"`
	DBName              string `mapstructure:"dbname"`
	Charset             string `mapstructure:"charset"`
	Prefix              string `mapstructure:"prefix"`
	Port                int    `mapstructure:"port"`
	MaxOpenConns        int    `mapstructure:"max_open_conns"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns"`
	LogMode             string `mapstructure:"log_mode"`
	LogFilename         string `mapstructure:"log_filename"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer"`
}
