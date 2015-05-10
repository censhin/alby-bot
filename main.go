package main

import (
    "net/http"

    "github.com/censhin/alby-bot/handlers"
    "github.com/censhin/alby-bot/slack"
)

func InitServer() {
    slack.ChatPostMessage("#bot-testing", "NERDS")
    slack.RtmConnect()
    handlers.InitHandlers()
    http.ListenAndServe(":8080", nil)
}

func main() {
    InitServer()
}
