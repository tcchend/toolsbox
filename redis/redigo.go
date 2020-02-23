package redis

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/garyburd/redigo/redis"
)

// 开启连接池
var RedisClient *redis.Pool

func ConnectRedis(){
	host := "redis.host"
	pwd := "redis.psw"
	MaxIdle := 100
	MaxActive := 1024
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host, redis.DialPassword(pwd))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

}

// 单独开启链接
func demo(){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	// do somethings...
}

func SetValue(key, value string) error{
	rcc := RedisClient.Get()
	defer rcc.Close()
	_, err := rcc.Do("SELECT", 0)
	if err != nil {
		return err
	}
	return nil
}

// 设置带过期时间
func SetValueWithExpire(key, value,expire string)error{
	rcc := RedisClient.Get()
	defer rcc.Close()
	_, err = rcc.Do("SET", key, value, "EX", expire)
	if err != nil {
		return err
	}
	return nil
}

func GetValue(key string) (string, error){
	rcc := RedisClient.Get()
	defer rcc.Close()
	value, err := redis.String(rcc.Do("GET", key))
	if err != nil {
		return "",err
	}
	return value ,nil
}


// 批量存储
func HSetValue(key, field, value string)error{
	rcc := RedisClient.Get()
	defer rcc.Close()
	//_, err = rcc.Do("HMSET", key, "name", name, "address", address)
	_, err = rcc.Do("HMSET", key, field, value)
	if err != nil {
		return err
	}
	return nil
}

func HGetValue(key, field string) (string, error){
	rcc := RedisClient.Get()
	defer rcc.Close()
	//res, err := redis.Strings(rcc.Do("HGETALL", v))
	//result[k] = map[string]string{res[0]: res[1], "displayName": v}
	//result = append(result, map[string]string{res[0]: res[1], "displayName": v, res[4]: res[5], res[6]: res[7]})
	value, err := redis.String(rcc.Do("HMGET", key,field))
	if err != nil {
		return "",err
	}
	return value ,nil
}

func IsExist(key string)bool{
	rcc := RedisClient.Get()
	defer rcc.Close()
	isKeyExist, err := redis.Bool(rcc.Do("EXISTS", key))
	if err != nil {
		return false
	} else {
		return isKeyExist
	}
}

func DelByKey(key string)bool{
	rcc := RedisClient.Get()
	defer rcc.Close()
	_, err = c.Do("DEL", "mykey")
	if err != nil {
		return false
	}
	return true
}

func Keys(key string)([]string,error){
	rcc := RedisClient.Get()
	defer rcc.Close()
	values, err := redis.Strings(rcc.Do("KEYS", key+"*"))
	if err != nil {
		return []string{},err
	}
	return values,nil
}

func List(key string,values []string)error{
	rcc := RedisClient.Get()
	defer rcc.Close()
	for _,value := range values{
		_, err = rcc.Do("LPUSH", key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func LRangeList(key string,start,end string) []string {
	res := make([]string, 0)
	rcc := RedisClient.Get()
	defer rcc.Close()
	values, _ := redis.Values(c.Do("lrange", key, start, end))
	for _, v := range values {
		res = append(res, string(v.([]byte)))
	}
	return res
}


// json redis
func JsonToRedis(){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}

func PipelineDemo(){
	//Send(commandName string, args ...interface{}) error
	//Flush() error
	//Receive() (reply interface{}, err error)
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")
	c.Flush()
	c.Receive() // reply from SET
	v, err = c.Receive() // reply from GET
}