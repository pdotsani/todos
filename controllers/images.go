package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/jmcvetta/neoism.v1"
	"log"
	"strings"
)

type ImageController struct {
	beego.Controller
}

type Image struct {
	alt, url string
}

func (c *ImageController) Post() {
	// DB connection (redirect to model request later: models/neo4jconnect.go)
	// db, err := neoism.Connect("http://neo4j:password@localhost:7474/db/data")
	// if err != nil {
	// 	log.Fatal("Error(neoism): Could not establish connection to db -- %s", err)
	// }

	// ROUTE LOGIC

	// 1: Get parameters from post request
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Ctx.Request.Body)
	str := buf.String()
	var img Image
	dec := json.NewDecoder(strings.NewReader(str))
	err := dec.Decode(&img)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("str: ", str)
	fmt.Println("img: ", img)

	// 2: Create data entry (node)
	// node, err := db.CreateNode(neoism.Props{
	// 	"alt": img.alt,
	// 	"url": img.url})
	// if err != nil {
	// 	log.Fatalln("Error(neoism): Node was not created -- ", err)
	// }
	// defer node.Delete()
	// node.AddLabel("Image")
}

func (c *ImageController) Get() {
	// DB connection (redirect to model request later: models/neo4jconnect.go)
	db, err := neoism.Connect("http://neo4j:password@localhost:7474/db/data")
	if err != nil {
		log.Fatal("Error(neoism): Could not establish connection to db --", err)
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
