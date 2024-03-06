package testdata

import "github.com/Soma-dev0808/blog_api/models"

// テスト用データ
var ArticleTestData = []models.Article{
	models.Article{
		ID: 1,
		Title: "firstPost",
		Contents: "This is my first post",
		UserName: "saki",
		NiceNum: 9,
	},
	models.Article{
		ID: 2,
		Title: "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum: 6,
	},
}



var CommentTestData = []models.Comment{
	models.Comment{
		ArticleID: 1,
		CommentID: 1,
		Message: "1st comment yeah",
	},
	models.Comment{
		ArticleID: 1,
		CommentID: 2,
		Message: "welcome",
	},
}