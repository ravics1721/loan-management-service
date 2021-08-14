package main

import (
	"github.com/ravics1721/loan-management-service/config"
	"github.com/ravics1721/loan-management-service/server"
)

func main() {
	config.Init()
	server.Init()
	server.Start()
}
