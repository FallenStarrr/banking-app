package repo

import (
	"github.com/FallenStarrr/banking-app/config"
	"context"
	"github.com/go-redis/redis"
	"fmt"
)


type AccRepo interface {
	    GetAcc()
			SetAcc()
			BlockAcc()
}


type Repo struct {
	red  *redis.Client
}

func NewRepo() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
})




return rdb
}