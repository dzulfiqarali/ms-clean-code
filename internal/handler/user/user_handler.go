package user

import (
	"github.com/guregu/null"
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

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    "0000",
		"responseMessage": "Success", // cast it to string before showing
		"name":            data.Name,
	})
	return
}

func (h UserHandler) ResolveListUser(c *gin.Context) {

	req := dto.UserListRequest{
		Nama:       null.StringFrom(c.Query("nama")),
		Alamat:     null.StringFrom(c.Query("alamat")),
		Pendidikan: null.StringFrom(c.Query("pendidikan")),
		Page:       null.StringFrom(c.Query("page")),
		Size:       null.StringFrom(c.Query("size")),
	}

	data := h.UserService.ResovleListUserByFilter(req)

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    "0000",
		"responseMessage": "Success", // cast it to string before showing
		"data":            data,
	})
	return
}
