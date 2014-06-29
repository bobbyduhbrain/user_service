package main

import(
  "github.com/bobbyduhbrain/user_service/models"
  "github.com/bobbyduhbrain/user_service/controllers"
  "github.com/go-martini/martini"
  "github.com/stretchr/gomniauth"
  "github.com/stretchr/signature"
  "github.com/codegangsta/martini-contrib/render"
  "net/http"
)

func main(){
  gomniauth.SetSecurityKey(signature.RandomKey(64))

  //*
  //  Controllers
  //*
  sessionsController := new(controllers.SessionsController)
  
  //* 
  //  Martini
  //*
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Directory: "templates",
    Layout:    "layout",
  }))
  m.Use(PopulateAppContext)
  m.Get("/", sessionsController.New)
  m.Run()
}

func PopulateAppContext(martiniContext martini.Context, w http.ResponseWriter, request *http.Request, renderer render.Render) {
  appContext := &models.AppContext{Request: request, Renderer: renderer, MartiniContext: martiniContext}
  martiniContext.Map(appContext)
}