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
	DisplayPage bool      `datastore:"displaypage" json:"displaypage" datastore_type:"Boolean" verbose_name:"Display this page"`
	Title       string    `datastore:"title" json:"title" datastore_type:"String" verbose_name:"Title"`
	URL         string    `datastore:"url" json:"url" datastore_type:"String" verbose_name:"URL"`
	PageOrder   int       `datastore:"pageorder" json:"pageorder" datastore_type:"Integer" verbose_name:"Page Order"`
	Content     string    `datastore:"content,noindex" json:"content" datastore_type:"Text" verbose_name:"URL"`
	Images      string    `datastore:"images,noindex" json:"images" datastore_type:"Text" verbose_name:"Images"`
	ExternalURL string    `datastore:"externalurl" json:"externalurl" datastore_type:"String" verbose_name:"Link to ..."`
	Update      time.Time `datastore:"update" json:"update" datastore_type:"DateTime" verbose_name:"Updated At"`
	Create      time.Time `datastore:"created" json:"created" datastore_type:"DateTime" verbose_name:"Created At"`
}

type Article struct {
	DisplayPage bool      `datastore:"displaypage" json:"displaypage" datastore_type:"Boolean" verbose_name:"Display this page"`
	Title       string    `datastore:"title" json:"title" datastore_type:"String" verbose_name:"Title"`
	URL         string    `datastore:"url" json:"url" datastore_type:"String" verbose_name:"URL"`
	PageOrder   int       `datastore:"pageorder" json:"pageorder" datastore_type:"Integer" verbose_name:"Page Order"`
	Content     string    `datastore:"content,noindex" json:"content" datastore_type:"Text" verbose_name:"URL"`
	Images      string    `datastore:"images,noindex" json:"images" datastore_type:"Text" verbose_name:"Images"`
	TagString   string    `datastore:"tagstring,noindex" json:"tagstring" datastore_type:"Text" verbose_name:"TagString"`
	Tags        []string  `datastore:"tags,noindex" json:"tags" datastore_type:"StringList" verbose_name:"Tags"`
	ExternalURL string    `datastore:"externalurl" json:"externalurl" datastore_type:"String" verbose_name:"Link to ..."`
	Update      time.Time `datastore:"update" json:"update" datastore_type:"DateTime" verbose_name:"Updated At"`
	Create      time.Time `datastore:"created" json:"created" datastore_type:"DateTime" verbose_name:"Created At"`
}
