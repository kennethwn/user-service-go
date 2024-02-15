package atom_user

import (
	"net/http"
	"strconv"
	atom_user "user_service/atom/user"

	"github.com/gin-gonic/gin"
)

func GetDataUserById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	data, err := atom_user.GetDataUserByIdUseCase(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"data":    data,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    data,
			"message": "Success",
		})
	}
}
