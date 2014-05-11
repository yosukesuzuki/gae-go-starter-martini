package webapp

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

func init() {
    m := martini.Classic()
    m.Use(render.Renderer(render.Options{
        Directory:  "templates",
        Layout:     "layout",
        Extensions: []string{".tmpl", ".html"},
        Charset:    "UTF-8",
    }))
    m.Get("/", index)
    http.Handle("/", m)
}

func index(r render.Render) {
    r.HTML(200, "index", "hello world")
}
