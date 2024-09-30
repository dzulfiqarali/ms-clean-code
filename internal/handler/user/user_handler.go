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

	shared.Success(
		c,
		data,
		shared.SetStatusCode(http.StatusCreated),
		shared.SetMessage("Successful"),
	)
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

	data := h.UserService.ResolveListUserByFilter(req)

	shared.Success(
		c,
		data,
		shared.SetStatusCode(http.StatusOK),
		shared.SetMessage("Successful"),
	)
	return
}

func (h UserHandler) ResolveUserByName(c *gin.Context) {

	req := dto.UserListRequest{
		Nama: null.StringFrom(c.Param("nama")),
	}

	if isNumeric(req.Nama.String) {
		shared.Failed(c, shared.CustomError(h.UserService.Error(shared.MakeError(errSvc.BadRequest))))
		c.Abort()
		return
	}

	data, err := h.UserService.ResolveUserByName(req)
	if err != nil {
		shared.Failed(c, shared.CustomError(h.UserService.Error(err)))
		c.Abort()
		return
	}

	shared.Success(
		c,
		data,
		shared.SetStatusCode(http.StatusOK),
		shared.SetMessage("Successful"),
	)
	return
}

func isNumeric(word string) bool {
	return regexp.MustCompile(`\d`).MatchString(word)
	// calling regexp.MustCompile() function to create the regular expression.
	// calling MatchString() function that returns a bool that
	// indicates whether a pattern is matched by the string.
}
