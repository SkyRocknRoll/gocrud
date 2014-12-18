package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/fzzy/radix/extra/pool"
)

var rpool *pool.Pool

func main() {
	rpool, _ = pool.NewPool("tcp", "localhost:6379", 2000)
	// if err != nil {
	// 	log.Fatalln("Failed to connect to Redis server!!!")
	// }

	http.HandleFunc("/", redisHandler)
	fmt.Println("Server listening on 8000")
	http.ListenAndServe(":8000", nil)

}

func redisHandler(rw http.ResponseWriter, req *http.Request) {
	r, redisErr := rpool.Get()

	defer rpool.CarefullyPut(r, &redisErr)

	data, err := r.Cmd("GET", "foo").Str()
	if err != nil {
		// io.WriteString(rw, err.Error())
		// log.Fatal(err)
		fmt.Println(err.Error())
	} else {
		io.WriteString(rw, data)
	}

}
