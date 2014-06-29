package main

import(
  "github.com/go-martini/martini"
  "github.com/stretchr/gomniauth"
  "github.com/stretchr/signature"
)

func main(){
  gomniauth.SetSecurityKey(signature.RandomKey(64))
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello poop!"
  })
  m.Run()
}