package services

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/jmcvetta/napping.v3"
	"gopkg.in/jmcvetta/neoism.v1"
	"os"
)

type (
	// Services contains common properties
	Service struct {
		NeoSession *neoism.Database
		UserID     string
	}
)

//$ export NEO4JURL=neo4j:Password@localhost:7474/db/data/
func (this *Service) Prepare() (err error) {
	neo4jUrl := os.Getenv("NEO4JURL")
	this.NeoSession, err = neoism.Connect(neo4jUrl)
	if err != nil {
		log.Error(err, service.UserID, "Service.Prepare")
		return err
	}

	return err
}
