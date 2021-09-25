package main

import (
	"os"
	"testing"

	spa "bitbucket.org/matchmove/platform-utilities-sgprefund-service/app"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/go-tools/env"
	"github.com/stretchr/testify/assert"
	"gopkg.in/matchmove/rest.v2"
)

func TestMain(m *testing.M) {
	var err error
	var app spa.App

	if app, err = spa.NewApplicationServer(); err != nil {
		panic(err.Error())
	}

	app.SetInfo(GetApplicationInformation())

	go func() {
		if err = spa.Run(&app); err != nil {
			panic(err.Error())
		}
	}()

	os.Exit(m.Run())
}

func TestMain_Main_NoEnvVars(t *testing.T) {
	env.CaptureEnvironmentVars()
	env.PurgeEnvironmentVars()

	defer func() {
		recovery := recover()
		assert.Contains(t, recovery, "Failed to initialize application server:")
		env.RestoreCapturedEnvironmentVars()
	}()

	main()
}

func TestMain_Main_EmptyEnvVar(t *testing.T) {
	env.CaptureEnvironmentVars()
	env.PurgeEnvironmentVars()

	os.Setenv(constant.EnvAppDomain, " ")
	os.Setenv(constant.EnvAppEnvironment, rest.ServerEnvTesting)
	os.Setenv(constant.EnvAppRefDocs, "http://0.0.0.0:81")
	os.Setenv(constant.EnvAppAccessLog, "fees.log")

	defer func() {
		recovery := recover()
		assert.Contains(t, recovery, "should not be empty")

		env.RestoreCapturedEnvironmentVars()
	}()

	main()
}

func TestMain_Main_DomainError(t *testing.T) {
	env.CaptureEnvironmentVars()
	env.PurgeEnvironmentVars()

	os.Setenv(constant.EnvAppDomain, "xxx")
	os.Setenv(constant.EnvAppEnvironment, rest.ServerEnvTesting)
	os.Setenv(constant.EnvAppRefDocs, "http://0.0.0.0:81")
	os.Setenv(constant.EnvAppAccessLog, "fees.log")

	defer func() {
		recovery := recover()
		assert.Contains(t, recovery, "invalid URI for request")
		env.RestoreCapturedEnvironmentVars()
	}()

	main()
}

func TestMain_GetApplicationInformation(t *testing.T) {
	name, _, _ := GetApplicationInformation()
	assert.Equal(t, name, "Fees Management Service")
}
