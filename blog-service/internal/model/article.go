package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	desc          string `json:"desc"`
	content       string `json:"content"`
	coverImageUrl string `json:"cover_image_url"`
	state         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
