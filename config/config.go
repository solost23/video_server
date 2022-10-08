package config

type ServerConfig struct {
	Version        string      `mapstructure:"version"`
	DebugMode      string      `mapstructure:"debug_mode"`
	Addr           string      `mapstructure:"addr"`
	Name           string      `mapsturcture:"name"`
	MysqlConfig    MysqlConfig `mapstructure:"mysql"`
	RedisConfig    RedisConfig `mapstructure:"redis"`
	MinioConfig    MinioConfig `mapstructure:"minio"`
	DeleteCronTime string      `mapstructure:"delete_cron_time"`
	JWTConfig      JWTConfig   `mapstructure:"jwt"`
	Md5Config      Md5Config   `mapstructure:"md5"`
	LogConfig      LogConfig   `mapstructure:"log"`
}

type MysqlConfig struct {
	Addr           string `mapstructure:"addr"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	DB             string `mapstructure:"db"`
	CasbinDB       string `mapstructure:"casbin_db"`
	Charset        string `mapstructure:"charset"`
	DeleteCronTime string `mapstructure:"delete_cron_time"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type MinioConfig struct {
	EndPoint        string `mapstructure:"end_point"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	SecretAccesskey string `mapstructure:"secret_access_key"`
	UserSSL         bool   `mapstructure:"user_ssl"`
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
