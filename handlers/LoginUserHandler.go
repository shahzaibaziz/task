package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/taskAPi"
	"github.com/taskAPi/gen/models"

	"github.com/taskAPi/gen/restapi/operations/users"
)

type loginUserImplements struct {
	rt runtime.Runtime
}

func (impl *loginUserImplements) Handle(params users.LoginUserParams) middleware.Responder {
	loginSuccessful:=impl.rt.Service().LoginUserService(models.LoginUserDefinition{Password: params.LoginUserBody.Password, Email: params.LoginUserBody.Email})
	if loginSuccessful== true{
		Token, err :=impl.rt.Service().GeneratingTheJWT(*params.LoginUserBody.Email)
		if err!=nil{
			err.Error()
		}

		return users.NewLoginUserOK().WithPayload(&models.LoginSuccessResponseDefinition{Token: Token})
	}
	return users.NewRegisterUserBadRequest()
}

func NewLoginUserHandler(rt runtime.Runtime) users.LoginUserHandler {
	return &loginUserImplements{
		rt: rt,
	}
}
