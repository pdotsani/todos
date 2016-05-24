package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/jmcvetta/neoism.v1"
	"log"
)

type Neo4jPool struct {
	beego.Controller
}

// https://github.com/astaxie/beego/issues/295
func (n *Neo4jPool) get() {
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		log.Fatal("Error(neoism): Could not establish connection to db -- %s", err)
	} else {
		fmt.Println("Successfully connected at localhost:7474")
		return db
	}
}
