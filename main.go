package main

import (
	"github.com/heroku/braiton-home/cmd/departmentsearch"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	departmentsearch.Run()
}
