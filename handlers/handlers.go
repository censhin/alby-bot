package handlers

import (
    "encoding/json"
    "log"
    "net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
    msg, err := json.Marshal("I have to go to work tomorrow :smith:")
    if err != nil {
        log.Fatal(err)
    }
    w.Write(msg)
}

func InitHandlers() {
    http.HandleFunc("/health", healthCheck)
}
