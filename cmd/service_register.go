package main

import (
	"consul-go/discovery"
	"fmt"
	"net/http"
)

func main() {
	//如果服务部署在本机，且Consul服务部署在Docker中，localhost要替换成本机在网络中的ip地址，因为Docker容器访问的localhost是容器而不是宿主机
	ip := "localhost"
	port := 15000
	discovery.ConsulRegister("test1", fmt.Sprintf("%s:%d", ip, port))

	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", "0.0.0.0", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("you are visiting health check api"))
}
