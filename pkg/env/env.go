package env

import (
	"bufio"
	"os"
	"strings"
)

const (
	Name           = "ddns-go"
	Environment    = "/etc/PG.conf"
	PanabitPathKey = "PGPATH"

	RunDir  = "/run"
	Pidfile = RunDir + "/ddns-go.pid"

	Ramdisk            = "/usr/ramdisk"
	ExtensionHome      = Ramdisk + "/app"
	ExtensionDir       = ExtensionHome + "/" + Name
	ExtensionBinaryDir = ExtensionDir + "/bin"
	DdnsGoBinary       = ExtensionBinaryDir + "/ddns-go"

	Certificates        = "/etc/ssl/certs/ca-certificates.crt"
	BundledCertificates = ExtensionDir + "/certs/ca-bundle.crt"

	ControlPanelHome         = Ramdisk + "/admin"
	ExtensionCgiDir          = ControlPanelHome + "/cgi-bin/App/" + Name
	ExtensionWebTemplatesDir = ExtensionCgiDir + "/template"
)

var (
	StorageHome         = "/usr/panabit"
	ExtensionStorageDir = StorageHome + "/app/" + Name

	ExtensionConfigurationStorageDir = ExtensionStorageDir + "/config"
	DdnsGoConfiguration              = ExtensionConfigurationStorageDir + "/ddns-go.yaml"
	DdnsGoDefaultConfiguration       = ExtensionConfigurationStorageDir + "/ddns-go.default.yaml"
	ManagerConfiguration             = ExtensionConfigurationStorageDir + "/manager.yaml"
	ManagerDefaultConfiguration      = ExtensionConfigurationStorageDir + "/manager.default.yaml"

	ExtensionCgiStorageDir          = ExtensionStorageDir + "/web/cgi"
	ExtensionWebTemplatesStorageDir = ExtensionStorageDir + "/web/template"
)

func Init() {
	v, err := find(PanabitPathKey)
	if err != nil || v == "" {
		return
	}
	StorageHome = v
	ExtensionStorageDir = StorageHome + "/app/" + Name
	ExtensionConfigurationStorageDir = ExtensionStorageDir + "/config"
	ExtensionCgiStorageDir = ExtensionStorageDir + "/web/cgi"
	ExtensionWebTemplatesStorageDir = ExtensionStorageDir + "/web/template"
}

func find(key string) (string, error) {
	f, err := os.Open(Environment)
	if err != nil {
		return "", err
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()
		if strings.HasPrefix(key, t) {
			if w := strings.Split(t, "="); len(w) == 2 {
				return w[1], nil
			} else {
				return "", nil
			}
		}
	}
	if err := sc.Err(); err != nil {
		return "", err
	}
	return "", err
}
