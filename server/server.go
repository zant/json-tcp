package main

import (
  "encoding/json"
  "fmt"
  "net"
  "time"

  "github.com/zant/json-tcp/common"
)

func main() {
  service := "localhost:1200"
  tcpAddr, err := net.ResolveTCPAddr("tcp", service)
  common.CheckError(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  common.CheckError(err)

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }

    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)

    var message common.Message
    decoder.Decode(&message)
    fmt.Println(message.String())
    time.Sleep(1000 * time.Nanosecond)
    encoder.Encode(message)

    conn.Close()
  }
}
