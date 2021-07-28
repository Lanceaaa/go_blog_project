package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

func (d *Dao) CountArticle(title string, state uint8) (int, error) {
	article := model.Article{Title: title, State: state}

	return article.Count(d.engine)
}

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)

	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(title, desc, coverImageUrl, createdBy string, state uint8) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		CoverImageUrl: coverImageUrl,
		State:         state,
		Model:         &model.Model{CreatedBy: createdBy},
	}

	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title, desc, coverImageUrl string, state uint8, modifiedBy string) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		CoverImageUrl: coverImageUrl,
		State:         state,
		Model:         &model.Model{ID: id, ModifiedBy: modifiedBy},
	}

	return article.Update(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}

	return article.Delete(d.engine)
}
