package main

import (
	"fmt"
	"github.com/bg-dao/bg-common/log"
	"github.com/bg-dao/bg-common/log/lager"
)

func main() {
	log.InitWithFile("log.yaml")

	for i := 0; i < 1; i++ {
		log.Infof("Hi %s, system is starting up ...", "paas-bot")
		log.Info("check-info", lager.Data{
			"info": "something",
		})

		log.Debug("check-info", lager.Data{
			"info": "something",
		})

		log.Warn("failed-to-do-somthing", lager.Data{
			"info": "something",
		})

		err := fmt.Errorf("this is an error")
		log.Error("failed-to-do-somthing", err)

		log.Info("shutting-down")
	}
}
