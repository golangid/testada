package redis

import (
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
)

type RedisSuite struct {
	suite.Suite
	Host     string
	Password string
	DB       int
	Client   *redis.Client
}

func (r *RedisSuite) SetupSuite() {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     r.Host,
		Password: r.Password,
		DB:       r.DB,
	})
}

func (r *RedisSuite) TearDownSuite() {
	r.Client.Close()
}
