package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// 获取标签总数
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 获取标签数据列表
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	// 假设服务端这边获取数据延迟
	// time.Sleep(time.Minute)
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

// 创建标签
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// 修改标签
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	// db = db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0)
	// db = db.Model(&Tag{}).Where("is_del = ?", 0)
	// return db.Update(t).Error
	if err := db.Model(t).Where("is_del = ?", 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

// 删除标签
func (t Tag) Delete(db *gorm.DB) error {
	// return db.Where("id = ? AND is_del = ?", t.ID, 0).Delete(&t).Error
	return db.Where("is_del = ?", 0).Delete(&t).Error
}
