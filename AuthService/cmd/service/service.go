package service

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/cmd/database"
	"github.com/cmd/entity"
	"github.com/cmd/env"
	"github.com/gin-gonic/gin"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Postgres struct {
	Postgresql *sql.DB
}

// type F struct{
// 	Database Postgres
// }

func (d *Postgres) Auth(c *gin.Context) {

	var loginPayload entity.LoginPayload

	if err := c.ShouldBindBodyWithJSON(&loginPayload); err != nil {
		return
	}

	user, err := database.GetUser(loginPayload.Username, d.Postgresql)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.StandardResp{Message: "error getting user from db", StatusCode: http.StatusInternalServerError})
		return
	}

	if user.Username != loginPayload.Username {
		c.JSON(http.StatusBadRequest, entity.StandardResp{Message: "Login in failed", StatusCode: http.StatusBadRequest})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginPayload.Password), []byte(user.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.StandardResp{Message: "Login in failed", StatusCode: http.StatusBadRequest})
		return
	}
	// d:=	make(map[string]int)

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 200).Unix(),
	}

	var secret_words string = env.GetString("JWT_KEY", "ArequestforalongtextmessageSearchresultsshowIfthisisyourintentpleaseclarifythecontextandwhatyouwantthetexttobeabout")
	claimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := claimsToken.SignedString(secret_words)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.StandardResp{Message: "error generating token", StatusCode: http.StatusInternalServerError})
		return
	}

	t := entity.TokenRes{
		Token: token,
	}

	c.JSON(http.StatusOK, t)

}
