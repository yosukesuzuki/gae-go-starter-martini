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
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"net/http"
	//    "github.com/PuerkitoBio/goquery"
)

func init() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".html"},
		Funcs: []template.FuncMap{
			{
				"gaID":    GetGaID,
				"fbAppID": GetFbAppID,
			},
		},
		Charset: "UTF-8",
	}))
	m.Group("/admin", AdminRoute)
	m.Get("/", Index)
	http.Handle("/", m)
}

func Index(r render.Render) {
	data := map[string]interface{}{
		"title":       "top",
		"description": "this is a starter app for GAE/Go",
		"body":        "hello world",
	}
	r.HTML(200, "index", data)
}
