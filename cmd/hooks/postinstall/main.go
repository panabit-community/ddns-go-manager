package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"
)

func main() {
	generateConfigIfNotFound()
	start()
}

func generateConfigIfNotFound() {
	if _, err := os.Stat(env.DdnsGoConfiguration); err == nil {
		return
	} else if !os.IsNotExist(err) {
		os.Exit(1)
	}
	if err := env.CopyFile(env.DdnsGoDefaultConfiguration, env.DdnsGoConfiguration, 0644); err != nil {
		os.Exit(1)
	}
}

func start() {
	p, err := filepath.Abs("./appctrl")
	if err != nil {
		os.Exit(1)
	}
	cmd := exec.Command(p, "start")
	if err := cmd.Start(); err != nil {
		os.Exit(1)
	}
}
