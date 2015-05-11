package main

import (
    "github.com/censhin/alby-bot/util"
    "github.com/swenson/slacker"
)

var config *util.Config = util.GetConfig()

func main() {
    bot, err := slacker.NewBot(config.Token)
    if err != nil {
        return
    }

    bot.RespondWith("nerds", func(user string, matchParts []string) string {
        return "nerd"
    })

    select {}
}
