package main

import (
	//    "log"
	//    "fmt"
	//    "time"
	//    "strings"
	//    "strconv"
	//    "appengine"
	//    "appengine/urlfetch"
	//    "appengine/memcache"
	//    "html/template"
	//    "net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	//    "github.com/PuerkitoBio/goquery"
)

// AdminRoute manages routing for Restful API
func AdminRoute(r martini.Router) {
	r.Group("/rest/:modelName", func(restRoutes martini.Router) {
		restRoutes.Get("", AdminGetList)
		restRoutes.Get("/:id", AdminGetEntity)
		restRoutes.Post("", AdminNewEntity)
		restRoutes.Put("/:id", AdminUpdateEntity)
		restRoutes.Delete("/:id", AdminDeleteEntity)
	})
	r.Get("/", AdminIndex)
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
func AdminNewEntity(params martini.Params, r render.Render) {
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
