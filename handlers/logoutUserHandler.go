package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/taskAPi"
	"github.com/taskAPi/gen/restapi/operations/users"
)

type logoutUserImplemepents struct{
	rt runtime.Runtime
}

func (impl *logoutUserImplemepents) Handle(params users.LogoutUserParams) middleware.Responder {
	impl.rt.Service().LogoutUserService(*params.LogoutUserBody.Email,*params.LogoutUserBody.Token)
	return users.NewUpdateUserOK()
}

func NewLogoutUserHandler(rt runtime.Runtime) users.LogoutUserHandler {
	return &logoutUserImplemepents{
		rt: rt,
	}
}
