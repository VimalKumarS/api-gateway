package service

import (
    "github.com/urfave/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a server.
func NewServer() *negroni.Negroni {
    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mux := mux.NewRouter()
    repo := &repoHandler{}
    initRoutes(mux, formatter, repo)
    n.UseHandler(mux)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, repo repository) {
    mx.HandleFunc("/handle/{route:[0-9a-zA-Z//]*}", gateWayHandler(formatter, repo)).Methods("GET", "POST")
    mx.HandleFunc("/service", postAddServiceHandler(formatter, repo)).Methods("POST")
}
