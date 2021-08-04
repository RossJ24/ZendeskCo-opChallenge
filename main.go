package main

import (
	"github.com/RossJ24/ZendeskCo-opChallenge/cli"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cli := cli.New()
	cli.Run()
}
