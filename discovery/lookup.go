package discovery

import (
	"fmt"
	api "github.com/hashicorp/consul/api"
	"strconv"
)

const (
	consulAgentAddress = "localhost:32500"
)

// 从consul中发现服务
func ConsulLookUp(serviceId string) {
	// 创建连接consul服务配置
	config := api.DefaultConfig()
	config.Address = consulAgentAddress
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 获取指定service
	services, _, err := client.Health().Service(serviceId, "", true, nil)
	if err == nil {
		for _, service := range services {
			address := service.Service.Address
			port := service.Service.Port
			fmt.Println(address + ":" + strconv.Itoa(port))
		}
	}

	//只获取健康的service
	//serviceHealthy, _, err := client.Health().Service("service337", "", true, nil)
	//if err == nil{
	//	fmt.Println(serviceHealthy[0].Service.Address)
	//}

}
