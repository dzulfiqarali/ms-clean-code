package user

import (
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	errSvc "github.com/ms-clean-code/internal/domain/error"
	"github.com/ms-clean-code/internal/domain/user/model/dto"
	"github.com/ms-clean-code/shared"
	"net/http"
	"regexp"
)

func (h UserHandler) InsertDataUser(c *gin.Context) {
	var request dto.RegistUserRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		shared.Failed(c, shared.CustomError(h.UserService.Error(shared.MakeError(errSvc.BadRequest))))
		c.Abort()
		return
	}

	data, err := h.UserService.RegistrationUser(request)
	if err != nil {
		shared.Failed(c, shared.CustomError(h.UserService.Error(err)))
		c.Abort()
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

func (h UserHandler) ResolveUserByName(c *gin.Context) {

	req := dto.UserListRequest{
		Nama: null.StringFrom(c.Param("nama")),
	}

	if isNumeric(req.Nama.String) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"responseCode":    "5000",
			"responseMessage": "interval server error", // cast it to string before showing
		})
		return
	}

	data, err := h.UserService.ResovleUserByName(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"responseCode":    "4000",
			"responseMessage": "not found", // cast it to string before showing
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    "0000",
		"responseMessage": "Success", // cast it to string before showing
		"data":            data,
	})
	return
}

func isNumeric(word string) bool {
	return regexp.MustCompile(`\d`).MatchString(word)
	// calling regexp.MustCompile() function to create the regular expression.
	// calling MatchString() function that returns a bool that
	// indicates whether a pattern is matched by the string.
}
