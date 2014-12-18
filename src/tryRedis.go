package main

import (
	"fmt"
	"log"

	redis "github.com/fzzy/radix/redis"
)

func main() {
	r, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalln("Failed to connect to Redis server!!!")
	}
	data, err := r.Cmd("GET", "foo").Str()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

}
