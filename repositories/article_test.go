package repositories_test

import (
	"testing"

	"github.com/Soma-dev0808/blog_api/models"
	"github.com/Soma-dev0808/blog_api/repositories"
	"github.com/Soma-dev0808/blog_api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)
func TestSelectArticle(t *testing.T) {
	expected := testdata.ArticleTestData[0]

	got, err := repositories.SelectArticleDetail(testDB, expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != expected.ID {
		t.Errorf("ID: got %d but want %d\n", got.ID, expected.ID)
	}

	if got.Title != expected.Title {
		t.Errorf("Title: got %s but want %s\n", got.Title, expected.Title)
	}

	if got.Contents != expected.Contents {
		t.Errorf("Contents: got %s but want %s\n", got.Contents, expected.Contents)
	}

	if got.UserName != expected.UserName {
		t.Errorf("UserName: got %s but want %s\n", got.UserName, expected.UserName)
	}

	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: got %d but want %d\n", got.NiceNum, expected.NiceNum)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected models.Article
	} {
		{
			testTitle: "subTest1",
			expected: testdata.ArticleTestData[0],
		},
		{
			testTitle: "subTest2",
			expected: testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: got %d but want %d\n", got.ID, test.expected.ID)
			}
		
			if got.Title != test.expected.Title {
				t.Errorf("Title: got %s but want %s\n", got.Title, test.expected.Title)
			}
		
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: got %s but want %s\n", got.Contents, test.expected.Contents)
			}
		
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: got %s but want %s\n", got.UserName, test.expected.UserName)
			}
		
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: got %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title: "insertTest",
		Contents: "test test",
		UserName: "saki",
	}

	beforeCount, err := repositories.SelectArticleListCount(testDB)
	if err != nil {
		t.Error(err)
	}

	_, err = repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}

	afterCount, err := repositories.SelectArticleListCount(testDB)
	if err != nil {
		t.Error(err)
	}

	if afterCount - beforeCount != 1 {
		t.Errorf("new article id is expected %d but got %d\n", beforeCount + 1, afterCount)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		if _, err := testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName); err != nil {
			t.Fatal("Fail to clean up test data", err)
		}

	})
}

func TestUpdateNiceNum(t *testing.T) {
	articleID := 1

	beforeArticle, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	afterArticle, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	if afterArticle.NiceNum - beforeArticle.NiceNum != 1 {
		t.Errorf("updated nice num is expected %d but got %d\n", beforeArticle.NiceNum + 1 , afterArticle.NiceNum)
	}

	t.Cleanup(func() {
		const sqlStr = `
			update articles
			set nice = ?
			where article_id = ?
		`
		_, err := testDB.Exec(sqlStr, beforeArticle.NiceNum, articleID)
		if err != nil {
			t.Fatal("Fail to clean up test data", err)
		}
	})
}