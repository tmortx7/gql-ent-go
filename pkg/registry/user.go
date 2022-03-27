package registry

import (
	"github.com/tmortx7/gql-ent-go/pkg/adapter/controller"
	"github.com/tmortx7/gql-ent-go/pkg/adapter/repository"
	"github.com/tmortx7/gql-ent-go/pkg/usecase/usecase"
)

func (r *registry) NewUserController() controller.User {
	repo := repository.NewUserRepository(r.client)
	u := usecase.NewUserUsecase(repo)

	return controller.NewUserController(u)
}
