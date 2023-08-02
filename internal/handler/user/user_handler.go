package user

import (
	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model/dto"
	"github.com/gin-gonic/gin"
)

func (uh Handler) InsertDataUser(c *gin.Context) {
	var request dto.RegistUserRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		var resp map[string]interface{}
		resp["responseCode"] = "0101"
		resp["responseMessage"] = "Invalid Request Parameter"

		c.JSON(400, resp)
		return
	}

	data, err := uh.UserService.RegistrationUser(request)
	if err != nil {
		var resp map[string]interface{}
		resp["responseCode"] = "0102"
		resp["responseMessage"] = "Failed"

		c.JSON(400, resp)
		return
	}

	var resp map[string]interface{}
	resp["responseCode"] = "0000"
	resp["responseMessage"] = "Success"
	resp["name"] = data.Name
	c.JSON(200, resp)
}
