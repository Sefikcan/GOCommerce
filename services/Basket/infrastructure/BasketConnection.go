package Connection

import (
	"basket/common/constants"
	"basket/common/logging"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client
var Log = logging.HandleLogging()

func Connect(){
	client := redis.NewClient(&redis.Options{
		Addr: constants.BASKETCONNSTRING,
		Password: constants.PASSWORD,
		DB: constants.DBTYPE,
	})

	_, err := client.Ping().Result()
	if err != nil {
		Log.Panic("Could not connect to the database")
	}

	RedisClient = client
}
