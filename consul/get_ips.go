package consul

import (
	"log"
	"net"
)

func (c *Consul) GetIPs() (r []*net.IP) {
	m := map[string]*net.IP{}
	for _, consul := range c.consuls {
		members, error := consul.Agent().Members(true)
		if error != nil {
			log.Print("Unable to get members list:", error)
		}
		for _, member := range members {
			addr := member.Addr
			if m[addr] == nil {
				m[addr] = &[]net.IP{net.ParseIP(addr)}[0]
			}
		}
	}
	for _, v := range m {
		r = append(r, v)
	}
	return
}
