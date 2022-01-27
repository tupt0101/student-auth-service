package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tupt0101/student-auth-service/src/http"
	"github.com/tupt0101/student-auth-service/src/repository/db"
	"github.com/tupt0101/student-auth-service/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":" + os.Getenv("APP_PORT"))
}
