package test

import (
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/jmcvetta/neoism.v1"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	_ "todos/routers"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestNeo4j(t *testing.T) {
	db, err := neoism.Connect("http://neo4j:password@localhost:7474/db/data")

	Convey("Subject: Test Neo4j Connectivity", t, func() {
		Convey("No Errors", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Db Connect Response", func() {
			So(db, ShouldNotEqual, nil)
		})
	})
}
