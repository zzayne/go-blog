package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Article struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	Name       string     `json:"name"`
	Categories []Category `gorm:"many2many:article_category;ForeignKey:ID;AssociationForeignKey:ID" json:"categories"`
}

func (m *Article) List(cateID int, pager Pager, isBackend bool) (articles []Article, err error) {
	var category Category
	offset := (pager.PageNo - 1) * pager.PageSize

	if DB.First(&category, cateID).Error != nil {

		return nil, errors.New("分类ID错误")
	}
	if cateID != 0 {
		var sql = `SELECT distinct(articles.id), articles.name, articles.browse_count, articles.comment_count, articles.collect_count,  
					articles.status, articles.created_at, articles.updated_at, articles.user_id, articles.last_user_id  
				FROM articles, article_category  
				WHERE articles.id = article_category.article_id   
				{statusSQL}       
				AND article_category.category_id = {categoryID} 
				AND articles.deleted_at IS NULL 
				ORDER BY {orderField} {orderASC}
				LIMIT {offset}, {pageSize}`
		sql = strings.Replace(sql, "{categoryID}", strconv.Itoa(cateID), -1)
		sql = strings.Replace(sql, "{orderField}", pager.OrderField, -1)
		sql = strings.Replace(sql, "{orderASC}", pager.OrderASC, -1)
		sql = strings.Replace(sql, "{offset}", strconv.Itoa(offset), -1)
		sql = strings.Replace(sql, "{pageSize}", strconv.Itoa(pager.PageSize), -1)

		if isBackend {
			sql = strings.Replace(sql, "{statusSQL}", " ", -1)
		} else {
			sql = strings.Replace(sql, "{statusSQL}", " AND (status = 1 OR status = 2)", -1)
		}
		if err = DB.Raw(sql).Scan(&articles).Error; err != nil {
			return nil, err
		}
	} else {
		orderStr := pager.OrderField + " " + pager.OrderASC
		if isBackend {
			//管理员查询话题列表时，会返回审核未通过的话题
			err = DB.Offset(offset).Limit(pager.PageSize).
				Order(orderStr).Find(&articles).Error
		} else {
			err = DB.Where("status = 1 OR status = 2").Offset(offset).Limit(pager.PageSize).Order(orderStr).Find(&articles).Error
		}
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(articles); i++ {
			if err = DB.Model(&articles[i]).Related(&articles[i].Categories, "categories").Error; err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
		}

	}

	return articles, err
}

func (m *Article) TotalCount() int {
	return getCount("ID>0")
}

func getCount(maps interface{}) (count int) {
	DB.Model(&Article{}).Where(maps).Count(&count)

	return count
}
