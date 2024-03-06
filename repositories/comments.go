package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Soma-dev0808/blog_api/models"
)


func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	newId, err := result.LastInsertId()

	if err != nil {
		return models.Comment{}, err
	}

	newComment := models.Comment{CommentID: int(newId), ArticleID: comment.ArticleID, Message: comment.Message}

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	commentArray := make([]models.Comment, 0)
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		
		if err != nil {
			fmt.Println(err)
		} else {
			commentArray = append(commentArray, comment)
		}
	}

	return commentArray, nil
}