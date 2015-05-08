package main

import (
    "fmt"
    "log"
    //"io/ioutil"
    "encoding/json"
    "net/http"
    "github.com/censhin/alby-bot/util"
)

var config *util.Config = util.GetConfig()

func main() {
    url := config.Endpoint + "/api.test?token=" + config.Token
    resp, err := http.Get(url)
    if err != nil {
        log.Panic("hepl")
    }
    defer resp.Body.Close()
    body, err := json.Marshal(resp.Body)
    if err != nil {
        log.Panic("halp")
    }
    fmt.Println(body)
}
