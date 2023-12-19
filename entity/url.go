package entity

import "net/url"

// Estrutura para armazenar informações sobre os servidores backend
type backendServer struct {
	URL *url.URL
}

// Estrutura para armazenar informações sobre o load balancer
type loadBalancer struct {
	servers []*backendServer
	next    int
}
