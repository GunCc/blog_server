package config

// 这个结构体就是映射 viper所解析的文件 其中 yaml 就是去获取 配置文件中的字段
type Server struct {
	// Mysql Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	AutoCode Autocode        `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	Zap      Zap             `mapstructure:"zap" json:"zap" yaml:"zap"`
	System   System          `yaml:"system"`
	Mysql    Mysql           `yaml:"mysql"`
	DBList   []SpecializedDB `yaml:"db-list"`
}
