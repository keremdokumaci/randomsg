package main

import "github.com/keremdokumaci/sqs-random-message-generator/app"

func main() {
	cli := app.NewCli()
	cli.Run()
}
