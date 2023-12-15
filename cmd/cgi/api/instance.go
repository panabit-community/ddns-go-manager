package main

import (
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/ddnsgo"
)

func startInstance() (int, any) {
	pid, err := ddnsgo.Start(ddnsgo.DefaultStartOpts)
	if err != nil {
		return 1, err
	}
	return 0, pid
}

func stopInstance() (int, any) {
	if err := ddnsgo.Stop(); err != nil {
		return 1, err
	}
	return 0, "ok"
}

func describeInstanceStatus() (int, any) {
	pid, err := ddnsgo.Status()
	if err != nil {
		return 1, err
	}
	if pid == 0 {
		return 0, "inactive"
	} else {
		return 0, "active"
	}
}
