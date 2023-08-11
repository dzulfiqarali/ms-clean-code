package external

import (
	"github.com/ms-clean-code/configs"
	"github.com/ms-clean-code/external/fakeapi"
)

// external : populate all domain implementation  external
type ExternalStruct struct {
	faImpl fakeapi.FakeApiImpl
}

// provide
func ProvideFakeApi(config *configs.Config) ExternalStruct {
	return ExternalStruct{
		faImpl: fakeapi.NewClientRequest(config),
	}
}
