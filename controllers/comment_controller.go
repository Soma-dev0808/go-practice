package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Soma-dev0808/blog_api/apperrors"
	"github.com/Soma-dev0808/blog_api/controllers/services"
	"github.com/Soma-dev0808/blog_api/models"
	"github.com/Soma-dev0808/blog_api/utils"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

// p289

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "fail to decode json")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	comment := reqComment

	newComment, err := c.service.PostCommentService(comment)
	if err != nil {
		fmt.Println(err)
		apperrors.ErrorHandler(w, req, err)
		return
	} 

	utils.SendSuccessPostJSONResponse(w, newComment)
}

func (c *CommentController) CommentListHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	commentList, err := c.service.GetCommentListService(articleID)
	if err != nil {
		fmt.Println(err)
		apperrors.ErrorHandler(w, req, err)
		return
	}

	utils.SendSuccessJSONResponse(w, commentList)
}
