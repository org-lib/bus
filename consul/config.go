package consul

type Info struct {
	ID                             string   `json:"id"`                             // 服务节点的名称
	Name                           string   `json:"name"`                           // 服务名称
	Port                           int      `json:"port"`                           // 服务端口
	Tags                           []string `json:"tags"`                           // tag，可以为空
	Address                        string   `json:"address"`                        // 服务 IP
	ConsulAddress                  string   `json:"consulAddress"`                  // 注册中心 consul 地址
	CheckPort                      int      `json:"checkPort"`                      // 健康检查端口
	CheckTimeout                   string   `json:"Timeout"`                        //
	CheckInterval                  string   `json:"checkInterval"`                  // 健康检查间隔
	DeregisterCriticalServiceAfter string   `json:"deregisterCriticalServiceAfter"` // check失败后30秒删除本服务，注销时间，相当于过期时间
	GRPC                                    // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
}
type ClientInfo struct {
	Name    string `json:"name"`    // 服务名称
	Tag     string `json:"tag"`     // tag，可以为空
	Address string `json:"address"` // 服务 IP
}
type GRPC struct {
	GIP      string `json:"gIp"`
	GPort    string `json:"gPort"`
	GService string `json:"gService"`
}
