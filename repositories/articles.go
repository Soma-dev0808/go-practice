package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Soma-dev0808/blog_api/models"
)

func SelectArticleListCount(db *sql.DB) (int, error) {
	const sqlStr = `
		select count(*) from articles;
	`
	var articlesCount int
	err := db.QueryRow(sqlStr).Scan(&articlesCount)
	if err != nil {
		return 0, nil
	}

	return articlesCount, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr =`
		select article_id, title, contents, username, nice, created_at
		from articles
		limit ? offset ?;
	`

	articleArray := make([]models.Article, 0)
	rows, err := db.Query(sqlStr, 5, (page - 1) * 5)
	if err != nil {
		return articleArray, err
	}

	defer rows.Close()

	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`

	row := db.QueryRow(sqlStr, articleID)

	var article models.Article
	var createdTime sql.NullTime

	if err := row.Err(); err != nil {
		fmt.Println(err)
		return article, err
	}

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return article, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, err
}

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error){
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values
		(?, ?, ?, 0, now());
	`

	res, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)

	newArticle := models.Article{
		Title: article.Title,
		Contents: article.Contents,
		UserName: article.UserName,
	}
	

	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	newId, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	newArticle.ID = int(newId)

	// 	スタック領域: 関数呼び出しに伴って自動的に割り当てられ、関数の実行が終了すると解放されるメモリ領域です。
	//			 	ローカル変数は通常、このスタック領域に割り当てられます。

	// ヒープ領域:  プログラムの実行中に動的に割り当てられるメモリ領域で、
	// 			  明示的に解放するか、ガーベジコレクタによって自動的に解放されるまで持続します。
	//			  長生きするデータや、関数の呼び出しを超えて存続する必要があるデータは、ヒープ領域に割り当てられます。

	// return &newId, nilでも良いが、ローカル変数のアドレスを外部から参照するのは良くないらしいので、
	// メモリにnewIdの値をidPtrが指すメモリアドレスに割り当てから、idPtrを返す

	// int64だとmodels.ArticleID int に合わないので、型変換してから返す
	// idPtr := new(int)
	// *idPtr = int(newId)

	// メモリにint64型を割り当て、返す方法
	// idPtr := new(int64)
	// *idPtr = newId

	return newArticle, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNiceNum = `
		update articles 
		set nice = ?
		where article_id = ?;
	`
	_, err = tx.Exec(sqlUpdateNiceNum, niceNum + 1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
