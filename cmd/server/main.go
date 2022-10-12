package main

import "github.com/matheusteixeira7/fc-go-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
