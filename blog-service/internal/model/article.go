package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

// 获取文章总数
func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 获取文章列表
func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var article []*Article
	var err error
	if pageOffset > 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	err = db.Where("is_del = ?", 0).Find(&article).Error
	if err != nil {
		return nil, err
	}

	return article, nil
}

// 创建文章
func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

// 更新文章
func (a Article) Update(db *gorm.DB, values interface{}) error {
	// db = db.Model(&Article{}).Where("is_del = ?", 0)
	// return db.Update(a).Error
	if err := db.Model(a).Where("is_del = ?", 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

// 删除文章
func (a Article) Delete(db *gorm.DB) error {
	return db.Where("is_del = ?", 0).Delete(&a).Error
}
