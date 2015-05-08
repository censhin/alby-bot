package main

import (
    "log"
    "net/http"
    "github.com/censhin/alby-bot/handlers"
    "github.com/Bowery/slack"
)

func slackAuth() {
    client := slack.NewClient("xoxb-4799782923-ZiWtbOfehjWGelqql3Aw2ZPh")
    err := client.SendMessage("#general", "NERDS!", "albybot")
    if err != nil {
        log.Fatal("Cannot authenticate to Slack.")
    }
}

func InitServer() {
    slackAuth()
    handlers.InitHandlers()
    http.ListenAndServe(":8080", nil)
}

func main() {
    InitServer()
}
