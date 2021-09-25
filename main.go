package main

import (
	rest "bitbucket.org/matchmove/integration-svc-aub/app"
	"bitbucket.org/matchmove/integration-svc-aub/modules/constant"
)

var (
	app rest.App
)

// main application entry point
func main() {
	var (
		err error
	)

	if app, err = rest.NewApplicationServer(); err != nil {
		panic(err.Error())
	}

	if err = app.Migrate(); err != nil {
		panic(err.Error())
	}

	if err = rest.Run(&app); err != nil {
		panic(err.Error())
	}
}

// GetApplicationInformation this holds the information about the app (version, realm, etc)
func GetApplicationInformation() (string, string, string) {
	return constant.Name, constant.Realm, constant.Version
}
