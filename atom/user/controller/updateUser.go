package atom_user

import (
	"net/http"
	"strconv"
	atom_user "user_service/atom/user"

	"github.com/gin-gonic/gin"
)

func UpdateUser(context *gin.Context) {
	var inputData atom_user.InputUserModel
	id, _ := strconv.Atoi(context.Param("id"))

	context.ShouldBindJSON(&inputData)

	data, err := atom_user.UpdateUserUseCase(id, inputData)
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
