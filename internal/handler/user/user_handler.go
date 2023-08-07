package user

import (
	"net/http"

	"bitbucket.org/bridce/ms-clean-code/internal/domain/user/model/dto"
	"github.com/gin-gonic/gin"
)

func (h UserHandler) InsertDataUser(c *gin.Context) {
	var request dto.RegistUserRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "0101",
			"responseMessage": "Failed Bind Request", // cast it to string before showing
		})
		return
	}

	data, err := h.UserService.RegistrationUser(request)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "0102",
			"responseMessage": "Failed", // cast it to string before showing
			"name":            data.Name,
		})
		return
	}

	defer func() {
		if err != nil {
			h.l.Error(err, "handler : InsertDataUser", "", nil, request, err, nil, nil)
		}
		h.l.Success(request, data, nil, nil, "handler : InsertDataUser", "nil", nil)
	}()

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    "0000",
		"responseMessage": "Success", // cast it to string before showing
		"name":            data.Name,
	})
	return
}
