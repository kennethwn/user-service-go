package atom_user

import (
	"net/http"

	atom_user "user_service/atom/user"

	"github.com/gin-gonic/gin"
)

func GetDataUserByName(context *gin.Context) {
	var inputData atom_user.InputUserModel
	context.ShouldBindJSON(&inputData)

	data, err := atom_user.GetDataUserByNameUseCase(inputData.Employee_name)
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
