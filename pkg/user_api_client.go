package api

import (
	"github.com/go-njn/gonjn030-api-client/internal"
	. "github.com/go-njn/gonjn030-api-client/shared"
)

func NewUserApiClient(config Config) UserApiClient {
	//demand that internal New creates correct implementation
	//goland:noinspection GoVarAndConstTypeMayBeOmitted
	var result UserApiClient = internal.NewUserApiClient(config)

	return result
}
