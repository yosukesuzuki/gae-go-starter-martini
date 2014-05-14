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

func adminRoute(r martini.Router) {
    r.Get("/", adminIndex)
}

func adminIndex(r render.Render) {
    data := map[string]interface{}{
        "title":"admin top",
        "description":"this is a starter app for GAE/Go",
        "body": "admin index",
    }
    r.HTML(200, "index", data)
}
