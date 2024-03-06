package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 全件取得
// func main () {
// 	dbUser := "docker"
// 	dbPassword := "docker"
// 	dbDatabase := "sampledb"
// 	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
// 	db, err := sql.Open("mysql",dbConn)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()

// 	const sqlStr = `
// 		select *
// 		from articles;
// 	`

// 	rows, err := db.Query(sqlStr)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer rows.Close()

// 	articleArray := make([]models.Article, 0)

// 	for rows.Next() {
// 		var article models.Article
// 		var createdTime sql.NullTime
// 		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

// 		if createdTime.Valid {
// 			article.CreatedAt = createdTime.Time
// 		}

// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			articleArray = append(articleArray, article)
// 		}
// 	}

// 	fmt.Printf("%+v\n", articleArray)
// }

// 取得ByID
// func main() {
// 	dbUser := "docker"
// 	dbPassword := "docker"
// 	dbDatabase := "sampledb"
// 	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
// 	db, err := sql.Open("mysql", dbConn)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()

//
// 	articleID := 0
// 	const sqlStr = `
// 		select *
// 		from articles
// 		where article_id = ?
// 	`

// 	row := db.QueryRow(sqlStr, articleID)

// 	if err := row.Err(); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	var article models.Article
// 	var createdTime sql.NullTime

// 	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	if createdTime.Valid {
// 		article.CreatedAt = createdTime.Time
// 	}

// 	fmt.Printf("%+v\n", article)

// }

// Insert
// func main () {
// 	dbUser := "docker"
// 	dbPassword := "docker"
// 	dbDatabase := "sampledb"
// 	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

// 	db, err := sql.Open("mysql", dbConn)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()

// 	article := models.Article {
// 		Title: "insert test",
// 		Contents: "Can I insert data correctly?",
// 		UserName: "saki",
// 	}

// 	const sqlStr = `
// 		insert into articles (title, contents, username, nice, created_at) values
// 		(?, ?, ?, 0, now())
// 	`

// 	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(result.LastInsertId())
// 	fmt.Println(result.RowsAffected())
// }

// Update
// Transaction
func main () {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tx, err	:= db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	article_id := 1
	sqlStr := `
		select nice
		from articles
		where article_id = ?;
	`
	row := tx.QueryRow(sqlStr, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, niceNum+1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	tx.Commit()
}