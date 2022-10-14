package consul

//consul agent -dev
import (
	"errors"
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"net/http"
	_ "net/http/pprof"
)

// RegisterServer 注册服务的IP和地址
func RegisterServer(info *Info) error {

	// 初始化参数检查
	if info.CheckPort == 0 {
		return errors.New(fmt.Sprintf("consul CheckPort error : %v", "端口不能为0"))
	}
	err := CheckIPAddr(info.Address)
	if err != nil {
		return err
	}
	if info.CheckTimeout == "" {
		info.CheckTimeout = "3s"
	}
	if info.CheckInterval == "" {
		info.CheckInterval = "5s"
	}
	if info.DeregisterCriticalServiceAfter == "" {
		info.DeregisterCriticalServiceAfter = "30s"
	}
	config := consulApi.DefaultConfig()
	config.Address = info.ConsulAddress
	client, err := consulApi.NewClient(config)
	if err != nil {
		return errors.New(fmt.Sprintf("consul client error : %v", err.Error()))
	}
	registration := new(consulApi.AgentServiceRegistration)
	registration.ID = info.ID           // 服务节点的名称
	registration.Name = info.Name       // 服务名称
	registration.Port = info.Port       // 服务端口
	registration.Tags = info.Tags       // tag，可以为空
	registration.Address = info.Address // 服务 IP

	checkPort := info.CheckPort
	registration.Check = &consulApi.AgentServiceCheck{ // 健康检查
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
		Timeout:                        info.CheckTimeout,
		Interval:                       info.CheckInterval,                  // 健康检查间隔
		DeregisterCriticalServiceAfter: info.DeregisterCriticalServiceAfter, //check失败后30秒删除本服务，注销时间，相当于过期时间
		// GRPC:     fmt.Sprintf("%v:%v/%v", info.GRPC.GIP, info.GRPC.GPort, info.GRPC.GService),// grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		return errors.New(fmt.Sprintf("register server error : %v", err.Error()))
	}

	http.HandleFunc("/check", Check)
	err = http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)
	if err != nil {
		return err
	}
	return nil
}

// SearchServer 获取 server 注册的 IP和地址
func SearchServer(info *ClientInfo) (map[string]string, error) {
	var err error
	err = CheckIPAddr(info.Address)
	if err != nil {
		return nil, err
	}

	var lastIndex uint64
	config := consulApi.DefaultConfig()
	config.Address = info.Address //consul server

	client, err := consulApi.NewClient(config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("api new client is failed, error: %v", err.Error()))
	}
	services, metainfo, err := client.Health().Service(info.Name, info.Tag, true, &consulApi.QueryOptions{
		WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("retrieving instances from Consul, error: %v", err.Error()))
	}
	lastIndex = metainfo.LastIndex
	adders := map[string]string{}
	for _, service := range services {
		adders[info.Name] = service.Service.Address
		adders[fmt.Sprintf("%vport", info.Name)] = fmt.Sprintf("%v", service.Service.Port)
	}
	return adders, nil
}
