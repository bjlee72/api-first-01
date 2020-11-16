package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/go-openapi/loads"

	healthReader "api-first-01/api/swagger/handlers/health/get"
	helloCreator "api-first-01/api/swagger/handlers/v1/hellos/create"
	helloReader "api-first-01/api/swagger/handlers/v1/hellos/get"
	helloLister "api-first-01/api/swagger/handlers/v1/hellos/list"
	helloStatusReader "api-first-01/api/swagger/handlers/v1/hellos/status/get"
	apis "api-first-01/api/swagger/restapi"
	ops "api-first-01/api/swagger/restapi/operations"
	"api-first-01/config"
	"api-first-01/storage/dummy"
	"api-first-01/utils/logger"
)

func main() {
	log.Print("Server initiated")

	environment := config.GetEnvironment()
	contextCreator := config.NewContextFactory(environment)
	storage := &dummy.Storage{}

	//
	// Handler implementations.
	//
	v1CreateHelloHandler := helloCreator.Dependency{
		ContextCreator: contextCreator,
		Storage:        storage,
	}.NewHandler()

	v1ReadHelloHandler := helloReader.Dependency{
		ContextCreator: contextCreator,
		Storage:        storage,
	}.NewHandler()

	v1ListHellosHandler := helloLister.Dependency{
		ContextCreator: contextCreator,
		Storage:        storage,
	}.NewHandler()

	v1ReadHelloStatusHandler := helloStatusReader.Dependency{
		ContextCreator: contextCreator,
		Storage:        storage,
	}.NewHandler()

	v1HealthReader := healthReader.Dependency{ContextCreator: contextCreator}.NewHandler()

	//
	// Link public handlers to endpoints.
	//

	spec := getSwaggerSpec(apis.FlatSwaggerJSON)
	handlers := ops.NewHelloAPI(spec)

	handlers.HealthV1HealthCheckHandler = v1HealthReader

	handlers.HelloV1CreateHelloHandler = v1CreateHelloHandler
	handlers.HelloV1ReadHelloHandler = v1ReadHelloHandler
	handlers.HelloV1ListHellosHandler = v1ListHellosHandler
	handlers.HelloV1ReadHelloStatusHandler = v1ReadHelloStatusHandler

	//
	// start servers - if you want to create multiple servers in one runtime,
	//   you can also do that by adding more server below. In that way,
	//   you can handle usecases such as providing public-facing APIs and
	//   internal-facing APIs at the same time. For the example, please see Morse.

	publicServer := apis.NewServer(handlers)
	publicServer.ConfigureAPI()
	publicServer.Port = getPortNumberFromEnv("PORT", 8080)
	publicServer.EnabledListeners = []string{"http"}

	type swaggerServer interface {
		Shutdown() error
		Serve() error
	}

	servers := []swaggerServer{publicServer}

	var wg sync.WaitGroup
	wg.Add(len(servers))
	for _, server := range servers {
		srv := server
		go func(wg *sync.WaitGroup) {
			defer func(wg *sync.WaitGroup) {
				_ = srv.Shutdown()
				wg.Done()
			}(wg)

			if err := srv.Serve(); err != nil {
				logger.Fatalf(nil, "couldn't start the server: %s", err.Error())
			}
		}(&wg)
	}
	wg.Wait()
}

func getPortNumberFromEnv(env string, defaultValue int) int {
	port := os.Getenv(env)
	if port != "" {
		pn, err := strconv.Atoi(port)
		if err != nil {
			logger.Fatalf(nil, "couldn't parse the server number string to int: %s", err)
		}
		// override default
		return pn
	}
	return defaultValue
}

func getSwaggerSpec(jsonSpec json.RawMessage) *loads.Document {
	spec, err := loads.Analyzed(jsonSpec, "")
	if err != nil {
		logger.Fatalf(nil, "cannot load analyzed swagger specification: %s", err)
	}
	return spec
}
