package resolve_domain

func TestResolve(t *testing.T) {
	domain := "www.baidu.com"
	resove(domain)
}

func resolve(domain string) {
	addr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
	}
	ip := addr.IP.String()
	fmt.Println("ip:" + ip)
}
