package services

import (
	"database/sql"
	"errors"

	"github.com/Soma-dev0808/blog_api/apperrors"
	"github.com/Soma-dev0808/blog_api/models"
	"github.com/Soma-dev0808/blog_api/repositories"
)

// リファクタ前
// func GetArticleListService(page int) ([]models.Article, error) {
// 	emptyArticleArray := make([]models.Article, 0)
// 	db, err := connectDB()
// 	if err != nil {
// 		return emptyArticleArray, err
// 	}

// 	articleList, err := repositories.SelectArticleList(db, page)
// 	if err != nil {
// 		return emptyArticleArray, err
// 	}

// 	return articleList, nil
// }

// リファクタ後
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	emptyArticleArray := make([]models.Article, 0)
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to fetch data")
		return emptyArticleArray, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

// func GetArticleService(articleID int) (models.Article, error) {
// 	db, err := connectDB()
// 	if err != nil {
// 		return models.Article{}, err
// 	}
// 	defer db.Close()

// 	article, err := repositories.SelectArticleDetail(db, articleID)
// 	if err != nil {
// 		return models.Article{}, err
// 	}

// 	commentList, err := repositories.SelectCommentList(db, articleID)
// 	if err != nil {
// 		return models.Article{}, err
// 	}

// 	article.CommentList = append(article.CommentList, commentList...)

// 	return article, nil
// }

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		// db.QueryRowの結果が０件の場合はsql.ErrNoRowsが返るようになっている
		// 入れ子になったエラーの中身を比較するには === ではなく、error.Is
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// func PostArticleService(article models.Article) (models.Article, error) {
// 	db, err := connectDB()
// 	if err != nil {
// 		return models.Article{}, err
// 	}
// 	defer db.Close()

// 	newArticle, err := repositories.InsertArticle(db, article)
// 	if err != nil {
// 		return models.Article{}, err
// 	}

// 	return newArticle, nil
// }

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert record")
		return models.Article{}, err
	}

	return newArticle, nil
}

// func UpdateNiceService(articleID int) error {
// 	db, err := connectDB()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	if err := repositories.UpdateNiceNum(db, articleID); err != nil {
// 		return err
// 	}

// 	return nil
// }

func (s *MyAppService) UpdateNiceService(articleID int) error {
	if err := repositories.UpdateNiceNum(s.db, articleID); err != nil {
		// db.QueryRowの結果が０件の場合はsql.ErrNoRowsが返るようになっている
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "failed to find target record")
			return err
		}
		err = apperrors.NoTargetData.Wrap(err, "failed to find target record")
		return err
	}

	return nil
}