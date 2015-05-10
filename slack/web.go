package slack

import (
    "net/url"
)

func ApiTest() ([]byte, error) {
    uri := config.Endpoint + "/api.test?buttes=dongues"
    resp, err := Get(uri)
    return resp, err
}

func AuthTest() ([]byte, error) {
    uri := config.Endpoint + "/auth.test?token=" + config.Token
    resp, err := Get(uri)
    return resp, err
}

func ChatPostMessage(channel string, text string) ([]byte, error) {
    uri := config.Endpoint + "/chat.postMessage"

    payload := url.Values{}
    payload.Set("token", config.Token)
    payload.Set("username", config.Username)
    payload.Set("channel", channel)
    payload.Set("text", text)

    resp, err := PostForm(uri, payload)
    return resp, err
}

