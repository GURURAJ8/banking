package main

import (
	"github.com/GURURAJ8/banking/app"
	"github.com/GURURAJ8/banking/logger"
)

func main(){
	logger.Info("Starting banking application...")
	app.Start()
}
