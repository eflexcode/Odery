package service

import (
	"net/http"
	"strconv"

	"github.com/cmd/database"
	"github.com/cmd/entity"
	"github.com/gin-gonic/gin"
)

type Repo struct {
	Db *database.Repo
}

func (r *Repo) GetNotifications(c *gin.Context) {

	var userId = c.Param("userId")
	var p = c.Query("page")
	var l = c.Query("limit")

	page, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.StandardResponse{Message: "NAN Page", StatusCode: http.StatusBadRequest})
		return
	}
	limit, err := strconv.Atoi(l)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.StandardResponse{Message: "NAN Limit", StatusCode: http.StatusBadRequest})
		return
	}

	nResults, err := database.GetNotificationsPagination(c, r.Db.DatabaseMongo, userId, int64(page), int64(limit))

	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.StandardResponse{Message: "Internal server error", StatusCode: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, nResults)

}

func (r *Repo) GetNotification(c *gin.Context) {

	var id = c.Param("id")

	notification, err := database.GetNotification(c, id, r.Db.DatabaseMongo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.StandardResponse{Message: "Internal server error", StatusCode: http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, notification)
}
