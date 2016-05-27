package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/jmcvetta/neoism.v1"
)

type (
	// Services contains common properties
	Service struct {
		neoSession *neoism.Database
	}
)

//$ export NEO4JURL=neo4j:Password@localhost:7474/db/data/
//func (this *Service) Prepare() (err error) {
//	neo4jUrl := os.Getenv("NEO4JURL")

func (service *Service) Prepare() (err error) {
	service.neoSession, err = neoism.Connect("http://neo4j:rpT2qu1IX%21%21yYP59D@localhost:7474/db/data/")
	if err != nil {
		beego.Error("DB Prep error:: ", err.Error())
		return
	}

	return
}

func (service *Service) QueryAllNodes(actor string) {

	stmt := `
        MATCH (a:Person)-[:ACTS_IN]->(b:Movie)
        WHERE a.name = {actorSub}
        RETURN a.name, type(rel), b.title
        `

	params := neoism.Props{"actorSub": actor}

	res := []struct {
		// `json:` tags matches column names in query
		A   string `json:"a.name"`
		Rel string `json:"type(rel)"`
		B   string `json:"b.title"`
	}{}

	cq := neoism.CypherQuery{
		Statement:  stmt,
		Parameters: params,
		Result:     &res,
	}

	err := service.neoSession.Cypher(&cq)
	beego.Error("query error:: ", err.Error())

	fmt.Printf("queryAllMoviesActedIn(%d)\n", len(res))
	for i, _ := range res {
		n := res[i]
		fmt.Printf("  Node[%d] %+v\n", i, n.A, n.Rel, n.B)
	}
}

//** AJAX SUPPORT
//AjaxResponse returns a standard ajax response.
//func (service *Service) AjaxResponse(resultCode int, resultString string, data interface{}) {
//	response := struct {
//		Result       int
//		ResultString string
//		ResultObject interface{}
//	}{
//		Result:       resultCode,
//		ResultString: resultString,
//		ResultObject: data,
//	}
//
//	service.Data["json"] = response
//	service.ServeJson()}
