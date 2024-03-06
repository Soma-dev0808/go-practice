package services

import (
	"github.com/Soma-dev0808/blog_api/apperrors"
	"github.com/Soma-dev0808/blog_api/models"
	"github.com/Soma-dev0808/blog_api/repositories"
)

// func GetCommentListService(articleID int) ([]models.Comment, error) {
// 	emptyCommentArray := make([]models.Comment, 0)
// 	db, err := connectDB()
// 	if err != nil {
// 		return emptyCommentArray, err
// 	}
// 	defer db.Close()

// 	commentArray, err := repositories.SelectCommentList(db, articleID)
// 	if err != nil {
// 		return emptyCommentArray, err
// 	}

// 	return commentArray, nil
// }

func (s *MyAppService) GetCommentListService(articleID int) ([]models.Comment, error) {
	emptyCommentArray := make([]models.Comment, 0)
	commentArray, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.NAData.Wrap(err, "failed to get records")
		return emptyCommentArray, err
	}

	return commentArray, nil
}

// func PostCommentService(comment models.Comment) (models.Comment, error){
// 	db, err := connectDB()
// 	if err != nil {
// 		return models.Comment{}, err
// 	}
// 	defer db.Close()

// 	newComment, err := repositories.InsertComment(db, comment)
// 	if err != nil {
// 		return models.Comment{}, err
// 	}

// 	return newComment, nil
// }

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error){
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.NoTargetData.Wrap(err, "no target record")
		return models.Comment{}, err
	}

	return newComment, nil
}