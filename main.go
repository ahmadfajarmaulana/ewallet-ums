package main

import (
	"ewallet-ums/cmd"
	"ewallet-ums/helpers"
)

func main() {
	helpers.SetupConfig()

	helpers.SetupLogger()

	helpers.SetupMySQL()

	go cmd.ServeGRPC()

	cmd.ServeHTTP()
}
