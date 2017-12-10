package config

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

// GetRedis return *redis.Client
func GetRedis(redisHost, redisTLS, redisPassword, redisPort string) (*redis.Client, error) {
	tlsSecured, err := strconv.ParseBool(redisTLS)

	if err != nil {
		return nil, err
	}

	var conf *tls.Config

	if tlsSecured {
		conf = &tls.Config{
			InsecureSkipVerify: tlsSecured,
		}
	} else {
		conf = nil
	}

	cl := redis.NewClient(&redis.Options{
		Addr:      fmt.Sprintf("%v:%v", redisHost, redisPort),
		Password:  redisPassword,
		DB:        0, // use default DB
		TLSConfig: conf,
	})
	return cl, nil
}
