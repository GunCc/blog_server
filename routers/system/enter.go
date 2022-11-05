package system

type RouterGroup struct {
	ApiRouter
	BaseRouter
	UserRouter
	ArticleTypeRouter
}
