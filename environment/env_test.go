package environment

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_GetListenPortDefaultsIfUnset(t *testing.T) {
	os.Unsetenv("PORT")

	assert.Equal(t, GetListenPort(), 8080)
}

func Test_GetListenPortGetsEnviormentVariable(t *testing.T) {
	os.Unsetenv("PORT")
	os.Setenv("PORT", "8081")

	assert.Equal(t, GetListenPort(), 8081)
}

func Test_GetListenPortGoesToDefaultssIfInvalidPortEnviormentVariable(t *testing.T) {
	os.Unsetenv("PORT")
	os.Setenv("PORT", "notaninteger")

	assert.Equal(t, GetListenPort(), 8080)
}

func TestGetRconConnectUrlGetsEnviornmentVariables(t *testing.T) {
	os.Unsetenv("RCONSERVER")
	os.Unsetenv("RCONPORT")

	os.Setenv("RCONSERVER", "mock")
	os.Setenv("RCONPORT", "9999")

	assert.Equal(t, GetRconConnectUrl(), "mock:9999")
}
