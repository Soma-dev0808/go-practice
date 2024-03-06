package api

import (
	"database/sql"
	"net/http"

	"github.com/Soma-dev0808/blog_api/controllers"
	"github.com/Soma-dev0808/blog_api/services"

	"github.com/gorilla/mux"
)

/*
memo:
	"github.com/Soma-dev0808/blog_api/controllers.services"で定義したinterfaceを元に
	NewArticleController, CommentController構造体を生成する。

	services.NewMyAppService(db)で生成した構造体は
	NewArticleController, CommentControllerのどちらも満たしている。（GetArticleServiceやPostCommentServiceを実装しているので）

	aCon、cConをNewRouterに渡し、ルーティングに対応させる
*/


func NewRouter(db *sql.DB) *mux.Router {
	// db *sql.dbを元にMyAppService構造体を生成する
	ser := services.NewMyAppService(db)
	// 1. サービス構造体 MyAppService(変数 ser) をもとに、ArticleController(変数 aCon) とCommentController(変数 cCon) を作成する
	// 2. 2 つのコントローラ構造体から、gorilla/mux のルータを作成する
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	// ルータを作り、パスーハンドラ関数の対応付けを登録する
	r := mux.NewRouter()

	// Article
	// GET
	r.HandleFunc("/", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)

	// POST
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)

	// UPDATE
	r.HandleFunc("/article/nice/{id:[0-9]+}", aCon.UpdateNiceHandler).Methods(http.MethodPut)

	// Comment
	// GET
	r.HandleFunc("/comment/{id:[0-9]+}", cCon.CommentListHandler).Methods(http.MethodGet)

	// POST
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)


	return  r
}