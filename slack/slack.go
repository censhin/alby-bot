package slack

import (
    "io/ioutil"
    "log"
    "net/http"
    "net/url"

    "github.com/censhin/alby-bot/util"
)

var config *util.Config = util.GetConfig()

func Get(uri string) ([]byte, error) {
    resp, err := http.Get(uri)
    if err != nil {
        log.Panic("GET to %s failed\n", uri)
        return nil, err
    } else {
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Panic("ioutil.ReadAll could not parse response\n")
            return nil, err
        } else {
            return body, nil
        }
    }
}

func PostForm(uri string, payload url.Values) ([]byte, error) {
    resp, err := http.PostForm(uri, payload)
    if err != nil {
        log.Panic("POST Form to %s failed\n", uri)
        return nil, err
    } else {
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Panic("ioutil.ReadAll could not parse response\n")
            return nil, err
        } else {
            return body, nil
        }
    }
}

