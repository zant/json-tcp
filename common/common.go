package common

import (
  "fmt"
  "os"
)

type Message struct {
  Channel string
  Body    string
}

func (m Message) String() string {
  s := "Message send to: " + m.Channel + "\nWith content: " + m.Body
  return s
}

func CheckError(err error) {
  if err != nil {
    fmt.Println("Fatal error ", err.Error())
    os.Exit(1)
  }
}
