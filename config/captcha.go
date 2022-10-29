package config

type Captcha struct {
	KeyLong   int     `json:"key-long" yaml:"key-long"`
	ImgWidth  int     `json:"img-width" yaml:"img-width"`
	ImgHeight int     `json:"img-height" yaml:"img-height"`
	MaxSkew   float64 `json:"max-skew" yaml:"max-skew"`
	DotCount  int     `json:"dot-count" yaml:"dot-count"`
}
