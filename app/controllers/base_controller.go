package controllers

import (
	"text/template"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)
var tmpl = template.Must(template.
	ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/home/index.html"))
