package main

import (
	"fmt"
	"github.com/iafoosball/users-service/usersgrpc/server/redis"
)

func main() {
	d, e := redis.HGET("chuj", "SIZE")
	if e != nil {
		fmt.Print("%v", e.Error())
	}

	fmt.Printf("I was here %v", d)
}