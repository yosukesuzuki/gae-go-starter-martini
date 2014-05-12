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
    "net/http"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
//    "github.com/PuerkitoBio/goquery"
)

var FB_APP_ID string = "551781848179574"
var GA_ID string = "UA-38221851-1"

func init() {
    m := martini.Classic()
    m.Use(render.Renderer(render.Options{
        Directory:  "templates",
        Layout:     "layout",
        Extensions: []string{".html"},
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
