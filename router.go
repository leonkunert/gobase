package main

import (
    "net/http"
    "os"
    "os/signal"
    "github.com/gorilla/mux"
    "syscall"
)

func Router () {
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, os.Interrupt)
    signal.Notify(sigchan, syscall.SIGTERM)

    // start the router
    go func() {
        r := mux.NewRouter()

        r.HandleFunc("/path/to/", HTTPFunc).Methods("GET")

        if err := http.ListenAndServe(":"+CFG.HttpPort, r); err != nil {
            log("ListenAndServer", err.Error())
        }
    }()

    <-sigchan
}