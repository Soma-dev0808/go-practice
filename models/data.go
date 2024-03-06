package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message: "This is a first message for article 1",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message: "This is a second message for article 1",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID: 1,
		Title: "Super article",
		Contents: "This is contents",
		UserName: "Test user",
		NiceNum: 1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt: time.Now(),
	}

	Article2 = Article{
		ID: 2,
		Title: "2 article",
		Contents: "This is contents 2",
		UserName: "Coleman",
		NiceNum: 2,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt: time.Now(),
	}
)