package ddnsgo

import (
	"os/exec"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"
)

type StartOpts struct {
	ConfigPath string
}

var DefaultStartOpts = StartOpts{
	ConfigPath: env.ExtensionConfigurationStorageDir + "/ddns-go.yaml",
}

func Start(opts StartOpts) (int, error) {
	if pid, err := Status(); err != nil {
		return 0, err
	} else if pid != 0 {
		return pid, nil
	}
	cmd := exec.Command(
		env.DdnsGoBinary,
		"-c", opts.ConfigPath,
	)
	if err := cmd.Start(); err != nil {
		return 0, err
	}
	env.WritePidfile(env.Pidfile, cmd.Process.Pid)
	return cmd.Process.Pid, nil
}

func Stop() error {
	return nil
}

func Status() (int, error) {
	pid, err := env.ReadPidfile(env.Pidfile)
	if err != nil {
		return 0, err
	}
	if exists, err := env.DescribeProcessExists(pid); err != nil {
		return 0, err
	} else if exists {
		return pid, nil
	} else {
		return 0, nil
	}
}
