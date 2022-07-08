package proxy

var proxies []Proxy

func Initialize() {
	proxies = []Proxy{}
}

func GetProxies() []Proxy {
	return proxies
}

func AddProxy(proxy Proxy) {
	proxies = append(proxies, proxy)
}
