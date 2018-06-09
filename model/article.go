package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zzayne/go-blog/utils"
)

type Article struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `sql:"index" json:"deletedAt"`
	Title       string     `json:"title"`
	Categories  []Category `gorm:"many2many:article_category;ForeignKey:ID;AssociationForeignKey:ID" json:"categories"`
	BrowseCount uint       `json:"browseCount"`
	ContentType int        `json:"contentType"`
	Content     string     `json:"content"`
	HTMLContent string     `json:"htmlContent"`
	Status      int        `json:"status"`
	UserID      uint       `json:"userID"`
	User        User       `json:"user"`
}

const (
	// ArticleVerifying 审核中
	ArticleVerifying = 1

	// ArticleVerifySuccess 审核通过
	ArticleVerifySuccess = 2

	// ArticleVerifyFail 审核未通过
	ArticleVerifyFail = 3
)

//List  ...
func (m *Article) List(cateID int, pager Pager, isBackend, noContent bool) (articles []Article, err error) {
	var category Category
	offset := (pager.PageNo - 1) * pager.PageSize

	if cateID != 0 {

		if DB.First(&category, cateID).Error != nil {
			return nil, errors.New("分类ID错误")
		}

		var sql = `SELECT distinct(articles.id), articles.title, articles.browse_count, articles.comment_count, articles.collect_count,  
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
		for i := 0; i < len(articles); i++ {
			articles[i].Categories = []Category{category}
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

	for i := 0; i < len(articles); i++ {
		if err := DB.Model(&articles[i]).Related(&articles[i].User, "users").Error; err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		if noContent {
			articles[i].Content = ""
			articles[i].HTMLContent = ""
		}
	}

	return articles, err
}

// Info ...
func (m *Article) Info(articleID int, isBackend bool, format string) (article Article, err error) {

	if err = DB.First(&article, articleID).Error; err != nil {
		return article, errors.New("文章ID错误")
	}
	if article.Status == ArticleVerifyFail {
		return article, errors.New("文章ID错误")
	}

	if !isBackend {
		article.BrowseCount++
		if err := DB.Save(&article).Error; err != nil {
			return article, err
		}
	}
	if err := DB.Model(&article).Related(&article.User, "users").Error; err != nil {
		fmt.Println(err.Error())
		return article, err
	}

	if err := DB.Model(&article).Related(&article.Categories, "categories").Error; err != nil {
		fmt.Println(err.Error())
		return article, err
	}
	if format != "md" {
		if article.ContentType != ContentTypeMarkdown {
			article.HTMLContent = utils.AvoidXSS(article.HTMLContent)
		} else {
			article.HTMLContent = utils.MarkdownToHTML(article.Content)
		}
		article.Content = ""
	}

	return article, nil

}

//Save ...
func (m *Article) Save(uid uint, article Article, isEdit bool) error {

	var queryArticle Article
	if isEdit {
		if err := DB.First(&queryArticle, article.ID).Error; err != nil {
			return errors.New("无效的文章ID")
		}
	} else {
		article.UserID = uid
	}

	if isEdit {
		tempArticle := article
		article = queryArticle
		article.Title = tempArticle.Title
		if article.ContentType == ContentTypeHTML {
			article.HTMLContent = tempArticle.Content
		} else {
			article.Content = tempArticle.Content
		}
		article.Categories = tempArticle.Categories
	} else {
		article.BrowseCount = 0
		article.Status = ArticleVerifying
		article.ContentType = ContentTypeMarkdown
	}

	article.Title = strings.TrimSpace(article.Title)

	article.Content = strings.TrimSpace(article.Content)
	article.HTMLContent = strings.TrimSpace(article.HTMLContent)

	if article.HTMLContent != "" {
		article.HTMLContent = utils.AvoidXSS(article.HTMLContent)
	}

	if article.Title == "" {
		return errors.New("文章名称不能为空")
	}

	if article.Categories == nil || len(article.Categories) <= 0 {
		return errors.New("请选择板块")
	}

	for i := 0; i < len(article.Categories); i++ {
		var category Category
		if err := DB.First(&category, article.Categories[i].ID).Error; err != nil {
			return errors.New("无效的版块id")
		}
		article.Categories[i] = category
	}

	var saveErr error

	if isEdit {
		saveErr = DB.Save(&article).Error
	} else {
		saveErr = DB.Create(&article).Error

	}

	return saveErr
}

// TotalCount ...
func (m *Article) TotalCount() int {
	return getCount("ID>0")
}

func getCount(maps interface{}) (count int) {
	DB.Model(&Article{}).Where(maps).Count(&count)

	return count
}

// Delete 删除
func (m *Article) Delete(id int) error {
	var article Article

	if err := DB.First(&article, id).Error; err != nil {
		return errors.New("无效的ID")
	}

	err := DB.Delete(&article).Error
	return err

}
