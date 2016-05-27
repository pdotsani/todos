package test

import (
	"bytes"
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

func TestImages(t *testing.T) {
	// @POST
	img := []byte(`{"alt":"image","url":"http://someurl.net/image.jpeg"}`)
	r, _ := http.NewRequest("POST", "/api/images", bytes.NewBuffer(img))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("images", "TestImages", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test POST /api/images\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})

	// @GET
	r2, _ := http.NewRequest("GET", "/api/images", nil)
	w2 := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w2, r2)

	beego.Trace("images", "TestImages", "Code[%d]\n%s", w2.Code, w2.Body.String())

	Convey("Subject: Test GET /api/images\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestActors(t *testing.T) {
	// @POST
	actor := []byte(`{"William Shatner"}`)
	r, _ := http.NewRequest("POST", "/api/actors", bytes.NewBuffer(actor))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("actors", "TestActors", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test POST /api/actors\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})

	// @GET
	r2, _ := http.NewRequest("GET", "/api/actors", nil)
	w2 := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w2, r2)

	beego.Trace("actors", "ActorList", "Code[%d]\n%s", w2.Code, w2.Body.String())

	Convey("Subject: Test GET /actors\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
