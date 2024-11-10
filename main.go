package main

import (
	"github.com/sanatan01/hither/cli"
	"github.com/sanatan01/hither/pkg/log"
)

func main() {
	log.InitLogger()
	cli.Execute()
}