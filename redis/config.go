package redis

import "errors"

var (
	currConf = &redisConf{
		Endpoint: "localhost:6379",
		Password: "",
	}
	currDriver *redisDriver
)

type redisConf struct {
	Endpoint string
	Password string
}

func (c *redisConf) Update(conf *redisConf) error {
	if conf.Endpoint != "" {
		c.Endpoint = conf.Endpoint
	}
	if conf.Password != "" {
		c.Password = conf.Password
	}
	_, err := c.VerifyDriver()
	return err
}

func (c *redisConf) VerifyDriver() (*redisDriver, error) {
	driver, err := NewRedisDriver(&RedisOptions{
		Endpoint: c.Endpoint,
		Password: c.Password,
	})
	if err != nil {
		return nil, err
	}
	if driver.Ping() == nil {
		currDriver = driver
		return currDriver, nil
	}

	return nil, errors.New("连接失败")
}
