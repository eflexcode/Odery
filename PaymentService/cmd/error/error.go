package error

import(
	"evn"
	"service"
	"net/http"

	"github.com/cmd/database"
	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context){
	s := StandardResponse{
		Message: "Internal server error",
		Status:  http.StatusInternalServerError,
	}

	c.JSON(http.StatusInternalServerError, s)
	return
}