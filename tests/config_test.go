package tests

import (
	"os"
	"testing"

	"github.com/rohitkeshwani07/go-bootstrap/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigDefault(t *testing.T) {
	os.Unsetenv("CONFIG_PATH")
	cfg, err := config.LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, "8080", cfg.Server.Port)
}

func TestLoadConfigFromFile(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "config-*.yaml")
	assert.NoError(t, err)
	data := []byte("server:\n  port: '1234'\n")
	_, err = f.Write(data)
	assert.NoError(t, err)
	f.Close()

	os.Setenv("CONFIG_PATH", f.Name())
	defer os.Unsetenv("CONFIG_PATH")

	cfg, err := config.LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, "1234", cfg.Server.Port)
}
