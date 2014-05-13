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
    "html/template"
    "net/http"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
//    "github.com/PuerkitoBio/goquery"
)

//Change This to Your FB App ID
var FbAppID = "551781848179574"
//Change This to Your Google Analytics ID
var GaID = "UA-38221851-1"

func getGaID() string {
    return GaID
}
func getFbAppID() string {
    return FbAppID
}

func init() {
    m := martini.Classic()
    m.Use(render.Renderer(render.Options{
        Directory:  "templates",
        Layout:     "layout",
        Extensions: []string{".html"},
        Funcs: []template.FuncMap{
            {
                "gaID": getGaID,
                "fbAppID": getFbAppID,
            },
        },
        Charset:    "UTF-8",
    }))
    m.Get("/", index)
    http.Handle("/", m)
}

func index(r render.Render) {
    data := map[string]interface{}{
        "title":"top",
        "description":"this is a starter app for GAE/Go",
        "body": "hello world",
    }
    r.HTML(200, "index", data)
}
