package services

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/jmcvetta/neoism.v1"
	"os"
)

type (
	// Services contains common properties
	Service struct {
		db *neoism.Database
	}
)

//$ export NEO4JURL=neo4j:Password@localhost:7474/db/data/
func (this *Service) Prepare() (err error) {
	neo4jUrl := os.Getenv("NEO4JURL")
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"bg_mess.log"}`)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)

	this.db, err = neoism.Connect(neo4jUrl)
	if err != nil {
		log.Error("Service.Prepare")
		return err
	}

	return err
}
