package repositories_test

import (
	"testing"

	"github.com/Soma-dev0808/blog_api/models"
	"github.com/Soma-dev0808/blog_api/repositories"
	"github.com/Soma-dev0808/blog_api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	expectedCommentNum := len(testdata.CommentTestData)
	
	comments, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	if len(comments) != expectedCommentNum {
		t.Errorf("expected comment count is %d but got %d", expectedCommentNum,  len(comments))
	}
}

func TestInsertComment(t *testing.T) {
	articleID := 1
	testComment := "comment for testing"
	newComment := models.Comment{
		ArticleID: articleID,
		Message: testComment,
	}

	beforeCommentList, err := repositories.SelectCommentList(testDB, articleID) 

	_, err = repositories.InsertComment(testDB, newComment)
	if err != nil {
		t.Error(err)
	}

	afterCommentList, err := repositories.SelectCommentList(testDB, articleID) 

	if len(afterCommentList) - len(beforeCommentList) != 1 {
		t.Errorf("expected comment count after inserting is %d but got %d", len(beforeCommentList)+1, len(afterCommentList))
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where article_id = ? and message = ?
		`
		if _, err := testDB.Exec(sqlStr, articleID, testComment); err != nil {
			t.Fatal("Fail to clean up test data", err)
		}
	})
}

func TestCommentDetail(t *testing.T) {
	articleID := 1
	tests := []struct {
		testTitle string
		expected models.Comment
	} {
		{
			testTitle: "subTest1",
			expected: testdata.CommentTestData[0],
		},
		{
			testTitle: "subTest2",
			expected: testdata.CommentTestData[1],
		},
	}

	comments, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	if len(comments) < len(tests) {
		t.Fatalf("The number of comments is less than the number of tests")
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			var found bool

			for _, comment := range comments {
				if comment.CommentID == test.expected.CommentID {
					found = true
					if comment.ArticleID != test.expected.ArticleID || comment.Message != test.expected.Message {
						t.Errorf("Expected %+v but got %+v", test.expected, comment)
					}
				}
			}

			if !found {
				t.Errorf("Comment with ID %d not found", test.expected.CommentID)
			}
		})
	}
}