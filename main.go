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
