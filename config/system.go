package config

// 系统配置文件

type System struct {
	Env           string `yaml:"env"`
	Addr          int    `yaml:"addr"`
	DbType        string `yaml:"db-type"`
	OssType       string `yaml:"oss-type"`
	UseRedis      bool   `yaml:"use-redis"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}
