package ddnsgo

// https://github.com/jeessy2/ddns-go/blob/afb8c71c0370c59aec2d60b9d2ac72197397b6c1/config/config.go#L59

// DnsConfig 配置
type DnsConfig struct {
	Ipv4 struct {
		Enable bool
		// 获取IP类型 url/netInterface
		GetType      string
		URL          string
		NetInterface string
		Cmd          string
		Domains      []string
	}
	Ipv6 struct {
		Enable bool
		// 获取IP类型 url/netInterface
		GetType      string
		URL          string
		NetInterface string
		Cmd          string
		IPv6Reg      string // ipv6匹配正则表达式
		Domains      []string
	}
	DNS DNS
	TTL string
}

// DNS DNS配置
type DNS struct {
	// 名称。如：alidns,webhook
	Name   string
	ID     string
	Secret string
}

// User 登录用户
type User struct {
	Username string
	Password string
}

// Webhook Webhook
type Webhook struct {
	WebhookURL         string
	WebhookRequestBody string
	WebhookHeaders     string
}

type Config struct {
	DnsConf []DnsConfig
	User
	Webhook
	// 禁止公网访问
	NotAllowWanAccess bool
}
