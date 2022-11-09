package utils

var (
	ApiVerify   = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	LoginVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	// LoginVerify       = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify    = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	ArticleTypeVerify = Rules{"Title": {NotEmpty()}, "Icon": {NotEmpty()}}
	ArticleTagVerify  = Rules{"Content": {NotEmpty()}, "Color": {NotEmpty()}}
	IdVerify          = Rules{"ID": []string{NotEmpty()}}
	PageInfoVerify    = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	ArticleVerify     = Rules{"Title": {NotEmpty()}, "SysArticleType": {NotEmpty()}, "Tags": {NotEmpty()}, "Content": {NotEmpty()}}
)
