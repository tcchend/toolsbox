/**
 * redis
 * Created by tiger on Sun.Nov.2019
 */
package redis

import (
	"fmt"
	"time"

	"github.com/alecthomas/log4go"
	"github.com/go-redis/redis"
	"errors"
)

var redisClient *redis.Client

func NewRedisClient(host, pwd string, port, db, poolSize int) error {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: pwd,
		DB:       db,
		PoolSize: poolSize,
	})

	if _, err := client.Ping().Result(); err != nil {
		log4go.Error(err.Error())
		return err
	}

	redisClient = client
	return nil
}

func HIncrBy(key, field string, offset int) error {
	intCmd := redisClient.HIncrBy(key, field, int64(offset))
	if intCmd.Err() != nil {
		return intCmd.Err()
	}
	return nil
}

func Set(key, value string) (string, error) {
	statCmd := redisClient.Set(key, value, 0)
	res, err := statCmd.Result()
	if err != nil {
		log4go.Error("Set: ", err.Error())
		return "", err
	}

	return res, nil
}

func SetWithExpir(key, value string, expir time.Duration) (string, error) {
	statCmd := redisClient.Set(key, value, expir)
	res, err := statCmd.Result()
	if err != nil {
		log4go.Error("SetWithExpir: ", err.Error())
		return "", err
	}

	return res, nil
}

func Exist(key string) (int64, error) {
	boolCmd := redisClient.Exists(key)
	if boolCmd.Err() != nil {
		log4go.Error("HFieldExist: ", boolCmd.Err())
		return 0, boolCmd.Err()
	}
	return boolCmd.Val(), nil
}

func Get(key string) (string, error) {
	strCmd := redisClient.Get(key)
	if strCmd.Err() != nil {
		if strCmd.Err().Error() != "redis: nil" {
			log4go.Error("Get: ", strCmd.Err())
		}

		return "", strCmd.Err()
	}
	return strCmd.Val(), nil
}

func HSet(key, field, value string) (bool, error) {
	boolCmd := redisClient.HSet(key, field, value)
	if boolCmd.Err() != nil {
		log4go.Error("HSet: ", boolCmd.Err())
		return false, boolCmd.Err()
	}
	return boolCmd.Val(), nil
}

func HGet(key, field string) (string, error) {
	strCmd := redisClient.HGet(key, field)
	if strCmd.Err() != nil {
		if strCmd.Err().Error() != "redis: nil" {
			log4go.Error("HGet %s", strCmd.Err())
		}

		return "", strCmd.Err()
	}
	return strCmd.Val(), nil
}

func HKeys(key string) ([]string, error) {
	strSlice := redisClient.HKeys(key)
	if strSlice.Err() != nil {
		return nil, strSlice.Err()
	}
	return strSlice.Val(), nil
}

func ZAdd(key string, members ...redis.Z) error {
	intCmd := redisClient.ZAdd(key, members ...)
	if intCmd.Err() != nil {
		return intCmd.Err()
	}
	return nil
}

func ZRangeByScore(key string, opt redis.ZRangeBy) ([]string, error) {
	res := redisClient.ZRangeByScore(key, opt)
	if res.Err() != nil || res == nil {
		return nil, errors.New("redis: zrange failed")
	}
	return res.Val(), nil
}

func HMSet(key string, fields map[string]interface{}) (string, error) {
	statCmd := redisClient.HMSet(key, fields)
	if statCmd.Err() != nil {
		log4go.Error("HMSet: ", statCmd.Err())
		return "", statCmd.Err()
	}
	return statCmd.Val(), nil
}

func HFieldExist(key, field string) (bool, error) {
	boolCmd := redisClient.HExists(key, field)
	if boolCmd.Err() != nil {
		log4go.Error("HFieldExist: ", boolCmd.Err())
		return false, boolCmd.Err()
	}
	return boolCmd.Val(), nil
}

func Close() {
	redisClient.Close()
}

//del
func Del(key string) (int64, error) {
	boolCmd := redisClient.Del(key)
	if boolCmd.Err() != nil {
		log4go.Error("KeyDel: ", boolCmd.Err())
		return 0, boolCmd.Err()
	}
	return boolCmd.Val(), nil
}
