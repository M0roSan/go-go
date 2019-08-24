package main

import (
	"fmt"
	"encoding/json"
	"github.com/go-redis/redis"
)

// Author struct
type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
func main() {
	fmt.Println("Go Redis")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("value:", val)

	jsonObj, err := json.Marshal(Author{Name: "Elly", Age:25})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", jsonObj, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err = client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}