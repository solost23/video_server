package configs

type ServerConfig struct {
	Version          string        `mapstructure:"version"`
	DebugMode        string        `mapstructure:"debug_mode"`
	TimeLocation     string        `mapstructure:"time_location"`
	Addr             string        `mapstructure:"addr"`
	PrometheusEnable bool          `mapstructure:"prometheus_enable"`
	ConfigPath       string        `mapstructure:"config_path"`
	MysqlConfig      MysqlConfig   `mapstructure:"mysql"`
	RedisConfig      RedisConfig   `mapstructure:"redis"`
	DeleteCronTime   string        `mapstructure:"delete_cron_time"`
	JWTConfig        JWTConfig     `mapstructure:"jwt"`
	Md5Config        Md5Config     `mapstructure:"md5"`
	LogConfig        LogConfig     `mapstructure:"log"`
	ZincConfig       ZincConfig    `mapstructure:"zinc"`
	ConsulConfig     ConsulConf    `mapstructure:"consul"`
	StaticOSS        StaticOSSConf `mapstructure:"static-oss"`
	OSSSrvConfig     OSSSrvConf    `mapstructure:"oss"`
}

type MysqlConfig struct {
	Addr           string `mapstructure:"addr"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	DB             string `mapstructure:"db"`
	Charset        string `mapstructure:"charset"`
	DeleteCronTime string `mapstructure:"delete_cron_time"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	Key      string `mapstructure:"key"`
	Duration uint   `mapstructure:"duration"`
}

type Md5Config struct {
	Secret string `mapstructure:"secret"`
}

type LogConfig struct {
	RuntimePath string `mapstructure:"runtime_path"`
	TrackPath   string `mapstructure:"track_path"`
}

type ZincConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type ConsulConf struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type StaticOSSConf struct {
	Domain string `mapstructure:"domain"`
}

type OSSSrvConf struct {
	Name string `mapstructure:"name"`
}
