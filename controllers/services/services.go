package services

import "github.com/Soma-dev0808/blog_api/models"

// type MyAppServicer interface {
// 	GetArticleService(articleID int) (models.Article, error)
// 	GetArticleListService(page int) ([]models.Article, error)
// 	PostArticleService(article models.Article) (models.Article, error)
// 	UpdateNiceService(articleID int) (error)

// 	PostCommentService(comment models.Comment) (models.Comment, error)
// 	GetCommentListService(articleID int) ([]models.Comment, error)
// }

type ArticleServicer interface {
	GetArticleService(articleID int) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	PostArticleService(article models.Article) (models.Article, error)
	UpdateNiceService(articleID int) (error)
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
	GetCommentListService(articleID int) ([]models.Comment, error)
}