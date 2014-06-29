package controllers

import(
  "github.com/bobbyduhbrain/user_service/models"
)

type SessionsController struct {
}

func (controller *SessionsController) New(appContext *models.AppContext) {
  appContext.Renderer.HTML(200, "sessions/new", "poop")
}

