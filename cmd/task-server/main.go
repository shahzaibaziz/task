package main

import (
	"github.com/go-openapi/loads"
	"log"
	runtime "wancloudsV2"
	"wancloudsV2/gen/restapi"
	"wancloudsV2/handlers"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "2.0")
	if err != nil {
		log.Fatalln(err)
	}
	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)

	}

	api := handlers.NewHandler(rt, swaggerSpec)
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}
	server.TLSCertificate = "../../crts/localhost.crt"
	server.TLSCertificateKey = "../../crts/localhost.key"
	server.Port = 8000
	server.Host = "127.0.0.1"

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}

}
