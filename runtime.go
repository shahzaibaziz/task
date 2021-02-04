package runtime

import (
	"log"

	"github.com/taskAPi/databaseLayer"
	"github.com/taskAPi/services"
)

type Runtime struct {
	svc *services.Service
}

func NewRuntime() (*Runtime, error) {
	conDb, err := databaseLayer.DatabaseConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Runtime{services.NewService(conDb)}, nil

	}

	func (r Runtime)Service()*services.Service{

		return  r.svc
	}
