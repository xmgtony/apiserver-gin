// Created on 2021/3/12.
// @author tony
// email xmgtony@gmail.com
// description

package ping

import "testing"

func TestIsLocalIp(t *testing.T) {
	ips := []struct {
		host   string
		expect bool
	}{
		{
			"192.168.1.1:9090",
			false,
		},
		{
			"127.0.0.1:8080",
			true,
		},
		{
			"localhost:8080",
			true,
		},
	}

	for _, ip := range ips {
		if isLocalIP(ip.host) == ip.expect {
			t.Logf("ip:host %s pass", ip.host)
		} else {
			t.Errorf("ip:host %s is not expected!", ip.host)
		}
	}
}
