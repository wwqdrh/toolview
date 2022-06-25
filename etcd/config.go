package etcd

import (
	"errors"
	"time"
)

// 饿汉单例 全局只维护这一个配置
var currConf = &etcdConf{
	Endpoints: []string{"127.0.0.1:2379"},
	UserName:  "",
	Password:  "",
}

var currDriver *EtcdDriver

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
	_, err := c.VerifyDriver()
	return err
}

func (c *etcdConf) VerifyDriver() (*EtcdDriver, error) {
	driver, err := NewEtcdDriver(&EtcdOptions{
		Endpoints:   currConf.Endpoints,
		DialTimeout: 3 * time.Second,
		OpTimeout:   3 * time.Second,
		UserName:    currConf.UserName,
		Password:    currConf.Password,
	})
	if err != nil {
		return nil, err
	}
	if driver.Ping() {
		currDriver = driver
		return currDriver, nil
	}
	return nil, errors.New("连接失败")
}
