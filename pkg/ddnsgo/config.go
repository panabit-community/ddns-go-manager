package ddnsgo

import (
	"os"
	"sync"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

	"gopkg.in/yaml.v3"
)

var (
	mutex sync.Mutex
)

type Config struct {
	User ConfigUser `yaml:"user"`
}

type ConfigUser struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func read() (*Config, error) {
	c := &Config{}
	bs, err := os.ReadFile(env.DdnsGoConfiguration)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(bs, c); err != nil {
		return nil, err
	}
	return c, nil
}

func readForUpdate() (*map[interface{}]interface{}, error) {
	m := make(map[interface{}]interface{})
	d, err := os.ReadFile(env.DdnsGoConfiguration)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(d, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func write(config *map[interface{}]interface{}) error {
	mutex.Lock()
	defer mutex.Unlock()
	d, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(env.DdnsGoConfiguration, d, 0644)
}

func Username() string {
	c, err := read()
	if err != nil {
		return ""
	}
	return c.User.Username
}

func SetUsername(username string) error {
	m, err := readForUpdate()
	if err != nil {
		return err
	}
	(*m)["user"].(map[interface{}]interface{})["username"] = username
	return write(m)
}

func Password() string {
	c, err := read()
	if err != nil {
		return ""
	}
	return c.User.Password
}

func SetPassword(password string) error {
	m, err := readForUpdate()
	if err != nil {
		return err
	}
	(*m)["user"].(map[interface{}]interface{})["password"] = password
	return write(m)
}
