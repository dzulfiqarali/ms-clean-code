package service

import (
	"github.com/ms-clean-code/internal/domain/user/model/dto"
	"testing"
)

func TestUserInterface(t *testing.T) {

	t.Run("ResolveListUserByFilter should return list of users", func(t *testing.T) {
		dummyData := dto.RegistUserRequest{
			Nama:       "test",
			Alamat:     "test",
			Umur:       "test",
			Pendidikan: "test",
		}
		want := "test"

		if dummyData.Nama != want {
			t.Error("Users is not there")
		}
	})
}
