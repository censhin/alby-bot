package slack

import (
    "encoding/json"
    "log"

    "golang.org/x/net/websocket"
)

func RtmStart() ([]byte, error) {
    uri := config.Endpoint + "/rtm.start?token=" + config.Token
    resp, err := Get(uri)
    return resp, err
}

func RtmConnect() (*websocket.Conn, error) {
    uri, err := RtmStart()
    if err != nil {
        log.Panic("RtmStart failed")
        return nil, err
    }

    b := make(map[string]interface{})
    json.Unmarshal(uri, &b)
    wsUrl := b["url"].(string)

    wsConn, err := websocket.Dial(wsUrl, "", "http://localhost")
    if err != nil {
        log.Panic("Cannot connect to websocket\n")
        return nil, err
    } else {
        return wsConn, nil
    }
}

