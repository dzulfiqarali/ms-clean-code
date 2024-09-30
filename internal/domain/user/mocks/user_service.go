package mocks

import "github.com/stretchr/testify/mock"

type MockUserInterface struct {
	mock.Mock
}

func (m *MockUserInterface) GetListUsers() (ret any) {
	return
}
