package model

import (
	"errors"
	"fmt"
	"time"
)

// Category 文章分类
type Category struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
	Name      string     `json:"name"`
	Sequence  int        `json:"sequence"` //同级别的分类可根据sequence的值来排序
	ParentID  int        `json:"parentId"` //直接父分类的ID
}

func (m *Category) List(pager Pager) (categories []Category, err error) {
	offset := (pager.PageNo - 1) * pager.PageSize
	err = DB.Offset(offset).Limit(pager.PageSize).Order("sequence asc").Find(&categories).Error
	return categories, err
}

func (m *Category) Save(cate Category, isNew bool) error {

	if isNew {
		//创建分类
		if err := DB.Create(&cate).Error; err != nil {
			return err
		}
	} else {
		var oldCate Category
		//更新分类
		if err := DB.First(&oldCate, cate.ID).Error; err == nil {
			updateMap := make(map[string]interface{})
			updateMap["name"] = cate.Name
			updateMap["sequence"] = cate.Sequence
			updateMap["parent_id"] = cate.ParentID
			if err := DB.Model(&oldCate).Updates(updateMap).Error; err != nil {
				fmt.Println(err.Error())
				return err

			}
		} else {
			return err
		}
	}

	return nil
}

func (m *Category) Find(id int) (cate Category, err error) {
	if err = DB.First(&cate, id).Error; err != nil {
		return cate, err
	}
	return cate, nil
}

func (m *Category) Delete(id int) error {
	var cate Category
	var childrenNum int
	if err := DB.First(&cate, id).Error; err != nil {
		return errors.New("无效的ID")
	}
	if err := DB.Model(&Category{}).Where("parent_id=?", id).Count(&childrenNum).Error; err != nil {
		return err
	}

	if childrenNum > 0 {
		return errors.New("该类别下存在子类别，无法删除")
	}

	if err := DB.Delete(&cate).Error; err != nil {
		return err
	}
	return nil

}
