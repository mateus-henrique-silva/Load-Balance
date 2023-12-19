package entity

import "net/url"

// Estrutura para armazenar informações sobre os servidores backend
type BackendServer struct {
	URL *url.URL
}

// Estrutura para armazenar informações sobre o load balancer
type LoadBalancer struct {
	Servers []*BackendServer
	Next    int
}
