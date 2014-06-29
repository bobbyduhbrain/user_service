package main

import(
  "github.com/go-martini/martini"
  "github.com/stretchr/gomniauth"
)

func main(){
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Run()
}