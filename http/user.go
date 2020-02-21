package http

import (
  "context"

  "github.com/ckbball/quik"
  "github.com/go-chi/chi"
)

type userHandler struct {
  router chi.Router
}
