package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/taskAPi"
	"github.com/taskAPi/gen/restapi/operations/users"
)

type updateUserHandler struct{
	rt runtime.Runtime
}

func (impl updateUserHandler) Handle(params users.UpdateUserParams) middleware.Responder {
	impl.rt.Service().UpdateUserService(*params.UpdateUserBody.JwtToken, *params.UpdateUserBody.Email,*params.UpdateUserBody.Name,*params.UpdateUserBody.Password )
	return users.NewUpdateUserOK()
}

func NewUpdateUserHandler(rt runtime.Runtime) users.UpdateUserHandler {
	return &updateUserHandler{
		rt: rt,
	}
}
