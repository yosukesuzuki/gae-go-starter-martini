package main

import (
	"log"
	//    "fmt"
	"reflect"
	"time"
	//    "strings"
	//    "strconv"
	"appengine"
	"appengine/datastore"
	//    "appengine/urlfetch"
	//    "appengine/memcache"
	//    "html/template"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	//    "github.com/PuerkitoBio/goquery"
)

// AdminRoute manages routing for Restful API
func AdminRoute(r martini.Router) {
	r.Get("/rest/metadata/:modelName", AdminGetMetaData)
	r.Group("/rest/:modelName", func(restRoutes martini.Router) {
		restRoutes.Get("", AdminGetList)
		restRoutes.Get("/:id", AdminGetEntity)
		restRoutes.Post("", AdminNewEntity)
		restRoutes.Put("/:id", AdminUpdateEntity)
		restRoutes.Delete("/:id", AdminDeleteEntity)
	})
	r.Get("/", AdminIndex)
}

type KindField struct {
	FieldName   string `json:"field_name"`
	FieldType   string `json:"field_type"`
	VerboseName string `json:"verbose_name"`
}

// AdminGetMetaData returns list of a kind
func AdminGetMetaData(params martini.Params, r render.Render) {
	var kind = &AdminPage{}
	var itemList []KindField
	s := reflect.ValueOf(kind).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		//f := s.Field(i)
		kindField := KindField{typeOfT.Field(i).Tag.Get("json"), typeOfT.Field(i).Tag.Get("datastore_type"), typeOfT.Field(i).Tag.Get("verbose_name")}
		itemList = append(itemList, kindField)
	}
	r.JSON(200, map[string]interface{}{"model_name": params["modelName"], "fields": itemList})
}

// AdminGetList returns list of a kind
func AdminGetList(params martini.Params, r render.Render) {
	r.JSON(200, map[string]interface{}{"id": params["id"], "model_name": params["modelName"]})
}

// AdminGetEntity get a entity of a kind
func AdminGetEntity(params martini.Params, r render.Render) {
	r.JSON(200, map[string]interface{}{"id": params["id"], "model_name": params["modelName"]})
}

// AdminNewEntity insert a entity to a kind
func AdminNewEntity(params martini.Params, r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	adminPage := &AdminPage{
		DisplayPage: true,
		Title:       "hoge",
		URL:         "hogeurl",
		Update:      time.Now(),
		Create:      time.Now(),
	}
	key := datastore.NewIncompleteKey(c, "AdminPage", nil)
	_, err := datastore.Put(c, key, adminPage)
	if err != nil {
		log.Println("test")
	}
	r.JSON(200, map[string]interface{}{"id": params["id"], "model_name": params["modelName"]})
}

// AdminUpdateEntity update a entity of a kind
func AdminUpdateEntity(params martini.Params, r render.Render) {
	r.JSON(200, map[string]interface{}{"id": params["id"], "model_name": params["modelName"]})
}

// AdminDeleteEntity delete a entity of a kind
func AdminDeleteEntity(params martini.Params, r render.Render) {
	r.JSON(200, map[string]interface{}{"id": params["id"], "model_name": params["modelName"]})
}

// AdminIndex renders index page for admin
func AdminIndex(r render.Render) {
	data := map[string]interface{}{
		"title":       "admin top",
		"description": "this is a starter app for GAE/Go",
		"body":        "admin index",
	}
	r.HTML(200, "index", data)
}
