/*
The SPRAWL server.

This app contains code to authenticate users, compute and serve game data.
Each path is defined in a separate file, (auth.go, api.go...).

Paths that require user authentication should be wrapped with the JwtVerifier
middleware. This middleware verifies the JWT that is passed with the request
and sets X-Username and X-Userid if the JWT checks out.

Game logic is applied under the /api endpoints Websocket handler (/api/websocket).

The general architecture of the game looks like this:

┌────┐                    ┌───────┐
│Game│─communicates with─►│clients│
└────┘                    └───────┘
  │
 has a
  │
  ▼
┌───┐
│ECS│
└───┘
*/
package main
import (
   _ "github.com/mattn/go-sqlite3"
   "github.com/peder2911/sparks/server/game"
   "net/http"
   "log"
)

var secret []byte = []byte("4321")

func main(){
   var err error 
   initialize(db_connect)
   mux := http.NewServeMux()
   game := game.NewGame()
   log.Println("Setting up")
   mux.Handle("/auth/", NewAuthHandler(secret, "/auth", db_connect))
   mux.Handle("/api/", NewJwtVerifier(NewApiHandler("/api", game), secret))
   go game.Loop()
   log.Println("Serving on :8000")
   err = http.ListenAndServe("0.0.0.0:8000", mux)
   panic(err)
}
