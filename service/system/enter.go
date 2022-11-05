package system

type SystemService struct {
	ApiService
	UserService
	JwtService
	OperationRecordService
	ArticleTypeService
	ArticleTagService
}
