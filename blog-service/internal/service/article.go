package service

type CountArticleRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"max:255"`
	IsDel         uint8  `form:"is_del" binding:"oneof=0 1"`
	State         uint8  `form:"state" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"required,min=3,max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,minx=3,max:255"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
	State         string `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=3,max=100"`
	Desc          string `form:"desc" binding:"required,min=3,max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,minx=3,max:255"`
	State         uint8  `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy    string `form:"modified_by" bindding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
