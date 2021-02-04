package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/taskAPi"
	"github.com/taskAPi/gen/models"
	"github.com/taskAPi/gen/restapi/operations/users"
)

type registerUserImpliments struct {
	rt *runtime.Runtime
}
func (impl *registerUserImpliments) Handle(params users.RegisterUserParams) middleware.Responder {
	//services.RegisterUserService(*params.RegisterUserBody.Name,*params.RegisterUserBody.Email, *params.RegisterUserBody.Password )

	isDataRedundent:=impl.rt.Service().VerifyData(models.RegisterUserDefinition{Email: params.RegisterUserBody.Email,Name: params.RegisterUserBody.Name,Password: params.RegisterUserBody.Password})
	if isDataRedundent==false{
		return users.NewRegisterUserBadRequest()
	}
	impl.rt.Service().RegisterUserService(models.RegisterUserDefinition{Name: params.RegisterUserBody.Name,Email: params.RegisterUserBody.Email, Password: params.RegisterUserBody.Password} )
	return users.NewRegisterUserOK()
}
func NewRegisterUserHandler(rt runtime.Runtime) users.RegisterUserHandler {
	return &registerUserImpliments{
		rt: &rt,
	}
}
