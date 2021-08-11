package main

import (
	"loan-management-service.com/lms/config"
	"loan-management-service.com/lms/server"
)

func main() {
	config.Init()
	server.Init()
	server.Start()
}
