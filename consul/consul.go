package consul

import (
	consulapi "github.com/hashicorp/consul/api"
	"strings"
)

type Consul struct {
	consuls []*consulapi.Client
}

func New(consuls []string, shouldRegister bool) (*Consul, error) {
	r := &Consul{}
	for _, oneConsul := range consuls {
		words := strings.Split(oneConsul, `/`)
		config := consulapi.DefaultConfig()
		config.Address = words[0]
		if len(words) > 1 {
			config.Datacenter = words[1]
		}
		consul, err := consulapi.NewClient(config)
		if err != nil {
			return nil, err
		}
		r.consuls = append(r.consuls, consul)
	}
	return r, nil
}
