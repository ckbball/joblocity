package main

import (
  "fmt"
  "os"

  "github.com/ckbball/quik"
  "github.com/ckbball/quik/db"
  "github.com/ckbball/quik/handler"
  "github.com/ckbball/quik/router"
  "github.com/ckbball/quik/store"
)

var (
  port = "8000"
)

func main() {
  if err := run(os.Args, os.Stdout); err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(exitFail)
  }
}

func run(args []string, stdout io.Writer) error {
  r := router.New()
  v1 := r.Group("/api")

  srvPort := port
  if os.Getenv("PORT") != "" {
    srvPort = os.Getenv("PORT")
  }
  addr := os.Getenv("LISTEN_ADDR")

  // grab db info needed here from env vars

  sqlDB := db.NewSql() // pass info here
  db.AutoMigrate(sqlDB)

  // start mongodb here
  datastore := db.NewMongo()

  us := store.NewUserStore(datastore)
  cs := store.NewCompanyStore(datastore)
  js := store.NewJobStore(sqlDB)
  as := store.NewApplicationStore(sqlDB)

  h := handler.NewHandler(us, cs, js, as)
  h.Register(v1)
  r.Logger.Fatal(r.Start(addr + ":" + srvPort))
}
