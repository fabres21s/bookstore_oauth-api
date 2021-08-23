package app

import (
	accesstoken "github.com/fabres21s/bookstore_oauth-api/src/domain/access_token"
	"github.com/fabres21s/bookstore_oauth-api/src/http"
	"github.com/fabres21s/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := accesstoken.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
