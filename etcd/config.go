package etcd

// 饿汉单例 全局只维护这一个配置
var currConf = &etcdConf{
	Endpoints: []string{"localhost:2379"},
}

type etcdConf struct {
	Endpoints []string `json:"endpoints"`
	UserName  string   `json:"username"`
	Password  string   `json:"password"`
}

func (c *etcdConf) Update(conf *etcdConf) error {
	if len(conf.Endpoints) > 0 {
		c.Endpoints = conf.Endpoints
	}
	if conf.UserName != "" {
		c.UserName = conf.UserName
	}
	if conf.Password != "" {
		c.Password = conf.Password
	}

	return nil
}
