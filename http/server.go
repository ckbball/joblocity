package http

import (
  "io"
  "io/ioutil"
  "net"
  "net/http"
  "net/url"
  "path"

  "github.com/ckbball/quik"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "golang.org/x/crypto/acme/autocert"
)

// Server represents an HTTP server.
type Server struct {
  ln net.Listener

  // Services
  UserService quik.UserService

  // Server options.
  Addr        string // bind address
  Host        string // external hostname
  Autocert    bool   // ACME autocert
  Recoverable bool   // panic recovery

  LogOutput io.Writer
}
