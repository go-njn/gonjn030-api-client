package shared

import (
	"context"
	"github.com/go-njn/gonjn030-api-client/pkg/domain"
)

type UserApiClient interface {
	GetAll(context.Context) ([]domain.User, error)
	GetById(context.Context, domain.ItemId) (domain.User, error)
	Create(context.Context, domain.User) (domain.ItemId, error)
	Update(context.Context, domain.ItemId, domain.User) error
	UpdateStatus(context.Context, domain.ItemId, domain.UserStatus) error
	UpdateGender(context.Context, domain.ItemId, domain.UserGender) error
	Delete(context.Context, domain.ItemId) error
}
