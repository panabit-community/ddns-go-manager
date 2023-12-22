package ddnsgo

import (
	"os"
	"sync"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/util"

	"gopkg.in/yaml.v3"
)

const (
	Address = "127.0.0.1:9876"
	Url     = "http://" + Address
)

var (
	mutex sync.Mutex
)

func Username() string {
	c := &Config{}
	if err := read(c); err != nil {
		return ""
	}
	return c.User.Username
}

func Password() string {
	c := &Config{}
	if err := read(c); err != nil {
		return ""
	}
	return c.User.Password
}

func GenerateCredentials() error {
	c := &Config{}
	if err := read(c); err != nil {
		return err
	}
	c.User.Username = util.RandomString(16)
	c.User.Password = util.RandomString(16)
	return persist(c)
}

func read(cfg *Config) error {
	f, err := os.ReadFile(env.DdnsGoConfiguration)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(f, cfg); err != nil {
		return err
	}
	return nil
}

func persist(cfg *Config) error {
	mutex.Lock()
	defer mutex.Unlock()
	d, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(env.DdnsGoConfiguration, d, 0644)
}
