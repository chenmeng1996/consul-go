package discovery

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	consulAddress = "localhost:32500"
)

func ConsulRegister(serviceName, endpoint string) {
	ipPort := strings.Split(endpoint, ":")
	localIp := ipPort[0]
	localPort, _ := strconv.Atoi(ipPort[1])

	// 创建连接consul服务配置
	config := api.DefaultConfig()
	config.Address = consulAddress
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	rand.Seed(time.Now().Unix())
	// 创建注册到consul的服务
	registration := new(api.AgentServiceRegistration)
	registration.ID = strconv.Itoa(rand.Int())
	registration.Name = serviceName
	registration.Port = localPort
	registration.Tags = []string{"testService"}
	registration.Address = localIp
	registration.EnableTagOverride = false

	// 增加consul健康检查回调函数
	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}

func ConsulDeregister(endpoint string) {
	// 创建连接consul服务配置
	config := api.DefaultConfig()
	config.Address = consulAddress
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	//取消注册
	deRegistration := new(api.CatalogDeregistration)
	deRegistration.Node = strings.Replace(endpoint, ":", "_", 1)

	_, _ = client.Catalog().Deregister(deRegistration, nil)
}
