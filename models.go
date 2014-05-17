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

// AdminPage store content for general pages
type AdminPage struct {
	DisplayPage bool      `datastore:"displaypage" json:"displaypage"`
	Title       string    `datastore:"title" json:"title"`
	URL         string    `datastore:"url" json:"url"`
	PageOrder   int       `datastore:"pageorder" json:"pageorder"`
	Content     string    `datastore:"content,noindex" json:"content"`
	Images      string    `datastore:"images,noindex" json:"images"`
	ExternalURL string    `datastore:"externalurl" json:"externalurl"`
	Update      time.Time `datastore:"update" json:"update"`
	Create      time.Time `datastore:"created" json:"created"`
}
