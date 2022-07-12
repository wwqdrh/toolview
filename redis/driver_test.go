package redis

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	goredis "github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestDriverSuite struct {
	suite.Suite

	isDev       bool
	redisClient *redisDriver
	redisMock   redismock.ClientMock
}

func TestRedisDriver(t *testing.T) {
	suite.Run(t, &TestDriverSuite{})
}

func (s *TestDriverSuite) SetupTest() {
	if os.Getenv("LOCAL") != "" {
		s.isDev = true
		driver, err := NewRedisDriver(&RedisOptions{
			Endpoint: os.Getenv("RedisEndpoint"),
			Password: os.Getenv("RedisPassword"),
		})
		require.Nil(s.T(), err)
		s.redisClient = driver
	} else {
		// mock client
		db, mock := redismock.NewClientMock()
		newsID := 123456789
		key := fmt.Sprintf("news_redis_cache_%d", newsID)

		mock.ExpectGet(key).RedisNil()
		mock.Regexp().ExpectSet(key, `[a-z]+`, 30*time.Minute).SetErr(errors.New("FAIL"))

		s.redisClient = &redisDriver{
			client: db,
		}
		s.redisMock = mock
	}
}

func (s *TestDriverSuite) TearDownTest() {
	if s.redisClient != nil {
		err := s.redisClient.client.Close()
		require.Nil(s.T(), err)
	}
	if s.redisMock != nil {
		err := s.redisMock.ExpectationsWereMet()
		require.Nil(s.T(), err)
	}
}

func (s *TestDriverSuite) TestStringCRUD() {
	if !s.isDev {
		s.T().Skip("not in local")
	}

	err := s.redisClient.client.Set(context.TODO(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := s.redisClient.client.Get(context.TODO(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := s.redisClient.client.Get(context.TODO(), "missing_key").Result()
	if err == goredis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}
}

func (s *TestDriverSuite) TestListCRUD() {
	if !s.isDev {
		s.T().Skip("not in local")
	}

	err := s.redisClient.client.RPush(context.TODO(), "queue", "message").Err()
	require.Nil(s.T(), err)

	result, err := s.redisClient.client.BLPop(context.TODO(), 1*time.Second, "queue").Result()
	require.Nil(s.T(), err)

	fmt.Println(result[0], result[1])
}

func (s *TestDriverSuite) TestHashCRUD() {
	if !s.isDev {
		s.T().Skip("not in local")
	}
}

func (s *TestDriverSuite) TestSortedSetCRUD() {
	if !s.isDev {
		s.T().Skip("not in local")
	}
}

func (s *TestDriverSuite) TestSetCRUD() {
	if !s.isDev {
		s.T().Skip("not in local")
	}
}
