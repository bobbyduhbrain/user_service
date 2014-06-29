package models

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
  "net/http"
)

type AppContext struct {
  Request        *http.Request
  Renderer       render.Render
  MartiniContext martini.Context
}