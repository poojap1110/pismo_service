package app_test

import (
	"fmt"
	"os"
	"testing"

	"bitbucket.org/matchmove/go-mock"
	"bitbucket.org/matchmove/go-tools/env"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/app"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"github.com/stretchr/testify/assert"
	"gopkg.in/matchmove/rest.v2"
)

func TestApp_RequiredEnvironmentVarNotDefined(t *testing.T) {
	var val string
	var err error

	assert := assert.New(t)

	env.CaptureEnvironmentVars(app.RequiredEnvironmentVars)
	env.PurgeEnvironmentVars(app.RequiredEnvironmentVars)

	_, err = app.RequiredEnvironmentVarNotDefined(constant.EnvAppDomain)
	assert.Error(err, "")

	env.RestoreCapturedEnvironmentVars()

	val, err = app.RequiredEnvironmentVarNotDefined(constant.EnvAppDomain)
	if err != nil {
		assert.Error(err, "")
		assert.Empty(val, "")
	}
}

func TestApp_NewApplicationServer_NoEnvDefined(t *testing.T) {
	env.CaptureEnvironmentVars()
	env.PurgeEnvironmentVars()

	_, err := app.NewApplicationServer()
	assert.Error(t, err, "")

	env.RestoreCapturedEnvironmentVars()
}

func TestApp_NewApplicationServer_EnvAppDomain(t *testing.T) {
	env.CaptureEnvironmentVars()

	os.Unsetenv(constant.EnvAppDomain)

	_, err := app.NewApplicationServer()
	assert.EqualError(t, err, fmt.Sprintf("Failed to initialize application server: %s is missing", constant.EnvAppDomain), "")

	env.RestoreCapturedEnvironmentVars()
}

func TestApp_NewApplicationServer_EnvAppEnvironment(t *testing.T) {
	env.CaptureEnvironmentVars()
	env.PurgeEnvironmentVars()

	loadApplicationEnvironment(t)

	os.Unsetenv(constant.EnvAppEnvironment)

	_, err := app.NewApplicationServer()
	assert.EqualError(t, err, fmt.Sprintf("Failed to initialize application server: %s is missing", constant.EnvAppEnvironment), "")

	env.RestoreCapturedEnvironmentVars()
}

func TestApp_NewApplicationServer_EnvAppRefDocs(t *testing.T) {
	env.CaptureEnvironmentVars()
	env.PurgeEnvironmentVars()

	loadApplicationEnvironment(t)

	os.Unsetenv(constant.EnvAppRefDocs)

	_, err := app.NewApplicationServer()
	assert.EqualError(t, err, fmt.Sprintf("Failed to initialize application server: %s is missing", constant.EnvAppRefDocs), "")

	env.RestoreCapturedEnvironmentVars()
}

func TestApp_NewApplicationServer_EnvAppAccessLog(t *testing.T) {
	env.CaptureEnvironmentVars()
	env.PurgeEnvironmentVars()

	loadApplicationEnvironment(t)

	os.Unsetenv(constant.EnvAppAccessLog)

	_, err := app.NewApplicationServer()
	assert.EqualError(t, err, fmt.Sprintf("Failed to initialize application server: %s is missing", constant.EnvAppAccessLog), "")

	env.RestoreCapturedEnvironmentVars()
}

func TestApp_NewApplicationServer(t *testing.T) {
	loadApplicationEnvironment(t)

	_, err := app.NewApplicationServer()
	assert.NoError(t, err, "")
}

func TestApp_GetApplication(t *testing.T) {
	_ = app.GetApplication()
}

func TestApp_SetGetRealm(t *testing.T) {
	loadApplicationEnvironment(t)

	a, err := app.NewApplicationServer()
	assert.NoError(t, err, "")

	a.SetRealm("fees")
	realm := a.GetRealm()

	assert.Equal(t, realm, "fees")
}

func TestApp_SetGetName(t *testing.T) {
	loadApplicationEnvironment(t)

	a, err := app.NewApplicationServer()
	assert.NoError(t, err, "")

	a.SetName("fees")
	name := a.GetName()
	assert.Equal(t, name, "fees")
}

func TestApp_SetGetVersion(t *testing.T) {
	loadApplicationEnvironment(t)

	a, err := app.NewApplicationServer()
	assert.NoError(t, err, "")

	a.SetVersion("1.0")
	version := a.GetVersion()
	assert.Equal(t, version, "1.0")
}

func TestApp_SetGetInfo(t *testing.T) {
	assert := assert.New(t)

	loadApplicationEnvironment(t)

	a, err := app.NewApplicationServer()
	assert.NoError(err, "")

	a.SetInfo("fees", "fees", "0.1")

	info := a.GetInfo()

	assert.Equal(info.Name, "fees")
	assert.Equal(info.Realm, "fees")
	assert.Equal(info.Version, "0.1")
}

func TestApp_GetWelcomeMessage(t *testing.T) {
	loadApplicationEnvironment(t)

	a, err := app.NewApplicationServer()
	assert.NoError(t, err, "")

	welcomeMessage := a.GetWelcomeMessage()
	assert.NotEmpty(t, welcomeMessage)
}

func TestApp_RooteHandler(t *testing.T) {
	loadApplicationEnvironment(t)

	a, err := app.NewApplicationServer()
	assert.NoError(t, err, "")

	a.SetInfo("fees", "fees", "0.1")

	w := mock.NewMockResponseWriter()
	app.RootHandler(w, nil)
}

func TestApp_Run(t *testing.T) {
	var err error
	var a app.App

	loadApplicationEnvironment(t)

	a, err = app.NewApplicationServer()
	assert.NoError(t, err, "")

	a.SetInfo(constant.Name, constant.Realm, constant.Version)

	go func() {
		defer func() {
			recovery := recover()
			t.Log(recovery)
		}()
		err = app.Run(&a)
		assert.NoError(t, err, "")
	}()
}

func TestApp_Run_NoEnv(t *testing.T) {
	var a app.App

	err := app.Run(&a)
	assert.Error(t, err, "")
}

func loadApplicationEnvironment(t *testing.T) {
	os.Setenv(constant.EnvAppDomain, "http://0.0.0.0:8080")
	os.Setenv(constant.EnvAppEnvironment, rest.ServerEnvTesting)
	os.Setenv(constant.EnvAppRefDocs, "http://0.0.0.0:81")
	os.Setenv(constant.EnvAppAccessLog, "platform-utilities-sgprefund-service.log")
}
