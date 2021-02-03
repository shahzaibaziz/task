package runtime

import (
	"log"
	"wancloudsV2/databaseLayer"
	"wancloudsV2/services"
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
