package main

import (
	"log"
    "encoding/json"
)

type Message struct {
    Type    string  `json:"type,omitempty"`
    From    string  `json:"from,omitempty"`
    To      string  `json:"to,omitempty"`
    Data    string  `json:"data,omitempty"`
}

func (m Message) String() string {

    data, err := json.Marshal(m)
    if err !=nil {
        log.Println(err)
        return ""
    }

    return string(data)
}