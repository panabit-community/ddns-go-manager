package env

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	mutex sync.Mutex
)

type Config struct {
	Global GlobalConfig `yaml:"global"`
}

type GlobalConfig struct {
	Enabled bool `yaml:"enabled"`
}

func GlobalEnabled() bool {
	c, err := read()
	if err != nil {
		return false
	}
	return c.Global.Enabled
}

func SetGlobalEnabled(enabled bool) error {
	c, err := read()
	if err != nil {
		return err
	}
	c.Global.Enabled = enabled
	return write(c)
}

func read() (*Config, error) {
	c := &Config{}
	bs, err := os.ReadFile(ManagerConfiguration)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(bs, c); err != nil {
		return nil, err
	}
	return c, nil
}

func write(config *Config) error {
	mutex.Lock()
	defer mutex.Unlock()
	d, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(ManagerConfiguration, d, 0644)
}
