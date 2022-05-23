package DaoRedis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)


//这里既可以传userId 也可以传friendId
func IsCelebrityRedis(userId string)(bool,error){

	RedisConnection, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return false,errors.New(fmt.Sprintf("redis.Dial err=", err))
	}
	defer RedisConnection.Close()

	//sismember
	celebrityKey := "celebrity"

	result, err :=redis.Bool(RedisConnection.Do("sismember",celebrityKey,userId))
	if err != nil {
		return false,errors.New(fmt.Sprintf("sismember celebrityKey error,err=", err))
	}
	return result,nil

}




