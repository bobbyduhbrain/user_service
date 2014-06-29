package main

import(
  "github.com/go-martini/martini"
  "github.com/stretchr/gomniauth"
  "github.com/stretchr/signature"
  "github.com/codegangsta/martini-contrib/render"
)

func main(){
  gomniauth.SetSecurityKey(signature.RandomKey(64))
  
  //* 
  //  Martini
  //*
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Directory: "templates",
    Layout:    "layout",
  }))
  m.Get("/", func() string {
    return "Hello poop!"
  })
  m.Run()
}