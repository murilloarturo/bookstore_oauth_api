package app

import (
	"github.com/gin-gonic/gin"
	"github.com/murilloarturo/bookstore_oauth_api/src/domain/access_token"
	"github.com/murilloarturo/bookstore_oauth_api/src/http"
	"github.com/murilloarturo/bookstore_oauth_api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:acces_token_id", atHandler.GetById)

	router.Run(":8080")
}
 