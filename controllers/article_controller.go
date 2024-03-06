package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Soma-dev0808/blog_api/apperrors"
	"github.com/Soma-dev0808/blog_api/controllers/services"
	"github.com/Soma-dev0808/blog_api/models"
	"github.com/Soma-dev0808/blog_api/utils"
	"github.com/gorilla/mux"
)

// p289

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Worlds!\n")
}

func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	fmt.Print(queryMap)

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		} 
	} else {
		page = 1
	}

	articles, err := c.service.GetArticleListService(page)
	if err != nil {
		fmt.Println(err)
		apperrors.ErrorHandler(w, req, err)
		return
	}

	utils.SendSuccessJSONResponse(w,articles)

	// マニュアルでのjsonエンコード
	// jsonData, err := json.Marshal(articles)

	// if err != nil {
	// 	errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
	// 	http.Error(w, errMsg, http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(jsonData)
}

func (c *ArticleController)ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "pathparam must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)

	if err != nil {
		fmt.Println(err)
		apperrors.ErrorHandler(w, req, err)
		return
	}

	utils.SendSuccessJSONResponse(w,article)
}

// POST
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	newArticle, err := c.service.PostArticleService(reqArticle)

	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	utils.SendSuccessJSONResponse(w, newArticle)

	// マニュアルでリクエストボディを取り出す
	// length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	// if err != nil {
	// 	http.Error(w, "cannot get content length\n", http.StatusBadRequest)
	// 	return
	// }
	// reqBodyBuffer := make([]byte, length)
	// if _, err := req.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
	// 	http.Error(w, "fail to get request body\n", http.StatusBadRequest)
	// 	return
	// }
	// defer req.Body.Close()

	// if err := json.Unmarshal(reqBodyBuffer, &reqArticle); err != nil {
	// 	http.Error(w,  "fail to decode json\n", http.StatusBadRequest)
	// 	return
	// }

	// jsonData, err := json.Marshal(article)

	// if err != nil {
	// 	http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
	// }

	// w.Write(jsonData)
}

func (c *ArticleController) UpdateNiceHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])

	if err != nil  {
		err = apperrors.BadParam.Wrap(err, "Invalid query parameter")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err := c.service.UpdateNiceService(articleID); err != nil {
		fmt.Println(err)
		apperrors.ErrorHandler(w, req, err)
		return
	}

	utils.SendSuccessUpdateJSONResponse(w)
}