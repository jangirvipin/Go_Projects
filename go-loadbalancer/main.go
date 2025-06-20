package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type LoadBalancer struct {
	Port            string
	Servers         []Server
	RoundRobinCount int
}

type SimpleServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

func (s *SimpleServer) Address() string {
	return s.Addr
}

func (s *SimpleServer) IsAlive() bool {
	//for simplicity, we assume all servers are alive
	return true
}

func (s *SimpleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Proxy.ServeHTTP(w, r)
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		Servers:         servers,
		RoundRobinCount: 0,
	}
}

func NewServer(addr string) *SimpleServer {
	serverURL, err := url.Parse(addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &SimpleServer{
		Addr:  addr,
		Proxy: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

func (lb *LoadBalancer) serveProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextServer()
	fmt.Println("targetServer:", targetServer)
	targetServer.ServeHTTP(w, r)
}

func (lb *LoadBalancer) getNextServer() Server {
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}
	lb.RoundRobinCount++
	return server
}

func main() {
	servers := []Server{
		NewServer("http://localhost:8081"),
		NewServer("http://localhost:8082"),
		NewServer("http://localhost:8083"),
	}

	lb := NewLoadBalancer("8080", servers)

	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	}

	http.HandleFunc("/", handleRedirect)
}
