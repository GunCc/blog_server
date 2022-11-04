package request

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

type Register struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	NickName     string `json:"nickname" gorm:"default:'mango'"`
	HeaderImg    string `json:"headerImg" gorm:"default:''"`
	AuthorityId  uint   `json:"authorityid" gorm:"default:888"`
	Enable       int    `json:"enable"`
	AuthorityIds []uint `json:"authorityIds"`
}
