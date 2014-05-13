package main

import (
//    "log"
//    "fmt"
    "time"
//    "strings"
//    "strconv"
//    "appengine"
//    "appengine/urlfetch"
//    "appengine/memcache"
//    "html/template"
//    "net/http"
//    "github.com/go-martini/martini"
//    "github.com/martini-contrib/render"
//    "github.com/PuerkitoBio/goquery"
)

type AdminPage struct {
        DisplayPage bool
        Title string
        Url string
        PageOrder int
        Content string
        Images string
        ExternalUrl string
        Update time.Time
        Create time.Time
}
