package config

type Server struct {
	// Mysql Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
}
