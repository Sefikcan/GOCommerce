package Connection

import (
	"basket/common/constants"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Connect(){
	client := redis.NewClient(&redis.Options{
		Addr: constants.BASKETCONNSTRING,
		Password: constants.PASSWORD,
		DB: constants.DBTYPE,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic("Could not connect to the database")
	}

	RedisClient = client
}
