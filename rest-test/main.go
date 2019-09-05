package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASS"), "rest_api_example")
	a.Run(":8080")
}