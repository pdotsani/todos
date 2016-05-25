package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gopkg.in/jmcvetta/neoism.v1"
	"log"
)

type ImageController struct {
	beego.Controller
}

type Image struct {
	alt, url string
}

func (c *ImageController) Post() {
	// DB connection (redirect to model request later: models/neo4jconnect.go)
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		log.Fatal("Error(neoism): Could not establish connection to db -- %s", err)
	}

	// ROUTE LOGIC

	// 1: Get parameters from post request
	var img Image
	jsoninfo := []byte(c.GetString("jsoninfo"))
	if err := json.Unmarshal(jsoninfo, &img); err != nil {
		log.Fatalln("Error(beego): Unmarshal Failed -- %s", err)
	}

	// 2: Create data entry (node)
	node, err := db.CreateNode(neoism.Props{
		"alt": img.alt,
		"url": img.url})
	if err != nil {
		log.Fatalln("Error(neoism): Node was not created -- %s", err)
	}
	defer node.Delete()
	node.AddLabel("Image")

	// 3: Successful Post send 201 (remove if beego does this by default)
	c.Ctx.Output.SetStatus(201)
}

func (c *ImageController) Get() {
	// DB connection (redirect to model request later: models/neo4jconnect.go)
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		log.Fatal("Error(neoism): Could not establish connection to db -- %s", err)
	}

	res := []struct {
		URL string `json:"a.url"`
		ALT string `json:"a.alt"`
		ID  string `json:"ID(n)"`
	}{}

	getAllImages := neoism.CypherQuery{
		Statement: `
			MATCH (i:Image)
			RETURN a.url, a.alt, ID(n)
		`,
		Result: &res,
	}

	db.Cypher(&getAllImages)

	result := res[0]
	c.Data["jsonp"] = &result
	c.ServeJSONP()
}
