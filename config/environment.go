package config

import (
	"log"
	"os"
)

// Environment represents the environment in which the server should run.
type Environment struct {
	port          string
	projectID     string
	projectNumber string
	region        string
	country       string
	stage         string
	datadogHostIP string
	imageVersion  string
}

// GetEnvironment returns Environment instance.
func GetEnvironment() *Environment {
	return &Environment{}
}

// ServiceName returns the service name.
func (env *Environment) ServiceName() string {
	return "api-first-01"
}

// Stage returns the stage value such as 'test' or 'prod'.
func (env *Environment) Stage() string {
	if env.stage == "" {
		env.stage = os.Getenv("STAGE")
		if env.stage == "" {
			env.stage = "test"
		}
		log.Printf("STAGE=%s", env.stage)
	}
	return env.stage
}

// Port returns the port number that morse should listen to.
func (env *Environment) Port() string {
	if env.port == "" {
		env.port = os.Getenv("PORT")
	}
	return env.port
}

// ProjectID returns the Google Cloud project ID.
func (env *Environment) ProjectID() string {
	if env.projectID == "" {
		env.projectID = os.Getenv("PROJECT_ID")
		log.Printf("PROJECT_ID=%s", env.projectID)
	}

	return env.projectID
}

// ProjectNumber returns the Google Cloud project number.
func (env *Environment) ProjectNumber() string {
	if env.projectNumber == "" {
		env.projectNumber = os.Getenv("PROJECT_NUMBER")
		log.Printf("PROJECT_NUMBER=%s", env.projectNumber)
	}

	return env.projectNumber
}

// DatadogHostIP returns the Datadog host IP address.
func (env *Environment) DatadogHostIP() string {
	if env.datadogHostIP == "" {
		env.datadogHostIP = os.Getenv("DATADOG_HOST_IP")
		log.Printf("DATADOG_HOST_IP=%s", env.datadogHostIP)
	}

	return env.datadogHostIP
}

// ImageVersion returns the image version of the service.
func (env *Environment) ImageVersion() string {
	if env.imageVersion == "" {
		env.imageVersion = os.Getenv("COMMIT")
		log.Printf("COMMIT(image version)=%s", env.imageVersion)
	}

	return env.imageVersion
}
