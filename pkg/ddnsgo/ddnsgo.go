package ddnsgo

import "xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

func Start() {}

func Status() (bool, error) {
	pid, err := env.ReadPidfile(env.Pidfile)
	if err != nil {
		return false, err
	}
	if s, err := env.DescribeProcessExists(pid); err != nil {
		return false, err
	} else {
		return s, nil
	}
}
